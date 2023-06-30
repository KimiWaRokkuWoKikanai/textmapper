#ifndef EXAMPLES_JSON_PARSER_H_
#define EXAMPLES_JSON_PARSER_H_

// generated by John; DO EDIT

#include <array>
#include <cstdint>
#include <ostream>
#include <string>
#include <utility>
#include <vector>

#include "absl/base/attributes.h"
#include "absl/functional/function_ref.h"
#include "absl/log/log.h"
#include "absl/status/status.h"
#include "absl/strings/str_format.h"
#include "lexer.h"

namespace json {

struct symbol {
  int32_t symbol = 0;
  int64_t offset = 0;
  int64_t endoffset = 0;
};

struct stackEntry {
  symbol sym;
  int8_t state = 0;
};

enum class NodeType {
  NoType,
  EmptyObject,
  JSONArray,
  JSONMember,
  JSONObject,
  JSONText,
  JSONValue,
  MultiLineComment,
  InvalidToken,
  JsonString,
  NonExistingType,
  NodeTypeMax
};
constexpr inline std::array<absl::string_view,
                            static_cast<size_t>(NodeType::NodeTypeMax)>
    nodeTypeStr = {
        "NONE",         "EmptyObject", "JSONArray",       "JSONMember",
        "JSONObject",   "JSONText",    "JSONValue",       "MultiLineComment",
        "InvalidToken", "JsonString",  "NonExistingType",
};
inline std::ostream& operator<<(std::ostream& os, NodeType t) {
  int i = static_cast<int>(t);
  if (i >= 0 && i < nodeTypeStr.size()) {
    return os << nodeTypeStr[i];
  }
  return os << "node(" << i << ")";
}

constexpr inline bool debugSyntax = true;
constexpr inline int startStackSize = 256;
constexpr inline int startTokenBufferSize = 16;
constexpr inline int32_t noToken = static_cast<int32_t>(Token::UNAVAILABLE);
constexpr inline int32_t eoiToken = static_cast<int32_t>(Token::EOI);

extern const int32_t tmAction[];
extern const NodeType tmRuleType[];
extern const int8_t tmRuleLen[];
extern const int32_t tmRuleSymbol[];

ABSL_MUST_USE_RESULT int8_t gotoState(int8_t state, int32_t symbol);
ABSL_MUST_USE_RESULT std::string symbolName(int32_t sym);

template <typename Lexer>
ABSL_MUST_USE_RESULT int32_t lookaheadNext(Lexer& lexer) {
  Token tok;
restart:
  tok = lexer.Next();
  switch (tok) {
    case Token::MULTILINECOMMENT:
    case Token::INVALID_TOKEN:
      goto restart;
    default:
      break;
  }
  return static_cast<int32_t>(tok);
}

ABSL_MUST_USE_RESULT int32_t lalr(int32_t action, int32_t next);

template <typename Lexer>
ABSL_MUST_USE_RESULT bool lookahead(Lexer& lexer_to_copy, int32_t next,
                                    int8_t start, int8_t end) {
  Lexer lexer = lexer_to_copy;
  std::vector<stackEntry> stack;
  stack.reserve(64);

  int8_t state = start;
  stack.push_back(stackEntry{.state = state});

  while (state != end) {
    int32_t action = tmAction[state];
    if (action < -2) {
      // Lookahead is needed.
      if (next == noToken) {
        next = lookaheadNext(lexer);
      }
      action = lalr(action, next);
    }

    if (action >= 0) {
      // Reduce.
      int32_t rule = action;
      auto ln = static_cast<int32_t>(tmRuleLen[rule]);

      stackEntry entry;
      entry.sym.symbol = tmRuleSymbol[rule];
      stack.resize(stack.size() - ln);
      state = gotoState(stack.back().state, entry.sym.symbol);
      entry.state = state;
      stack.push_back(std::move(entry));
    } else if (action == -1) {
      // Shift.
      if (next == noToken) {
        next = lookaheadNext(lexer);
      }
      state = gotoState(state, next);
      stack.push_back(stackEntry{
          .sym = symbol{.symbol = next},
          .state = state,
      });
      if (debugSyntax) {
        LOG(INFO) << "lookahead shift: " << symbolName(next) << " ("
                  << lexer.Text() << ")";
      }
      if (state != -1 && next != eoiToken) {
        next = noToken;
      }
    }

    if (action == -2 || state == -1) {
      break;
    }
  }

  if (debugSyntax) {
    LOG(INFO) << "lookahead done: " << ((state == end) ? "true" : "false");
  }

  return state == end;
}

template <typename Lexer>
ABSL_MUST_USE_RESULT bool AtEmptyObject(Lexer& lexer, int32_t next) {
  if (debugSyntax) {
    LOG(INFO) << "lookahead EmptyObject; next: " << symbolName(next);
  }
  return lookahead(lexer, next, 0, 42);
}

class Parser {
 public:
  template <typename Listener>
  explicit Parser(Listener&& listener)
      : listener_(std::forward<Listener>(listener)) {
    pending_symbols_.reserve(startTokenBufferSize);
  }

  template <typename Lexer>
  absl::Status Parse(Lexer& lexer) {
    return Parse(1, 44, lexer);
  }

 private:
  template <typename Lexer>
  absl::Status Parse(int8_t start, int8_t end, Lexer& lexer) {
    pending_symbols_.clear();
    int8_t state = start;
    std::vector<stackEntry> stack;
    stack.reserve(startStackSize);
    stack.push_back(stackEntry{.state = state});
    fetchNext(lexer, stack);
    while (state != end) {
      int32_t action = tmAction[state];
      if (action < -2) {
        // Lookahead is needed.
        if (next_symbol_.symbol == noToken) {
          fetchNext(lexer, stack);
        }
        action = lalr(action, next_symbol_.symbol);
      }
      if (action >= 0) {
        // Reduce.
        int32_t rule = action;
        int32_t ln = tmRuleLen[rule];
        stackEntry entry;
        entry.sym.symbol = tmRuleSymbol[rule];
        absl::Span<const stackEntry> rhs;

        if (ln == 0) {
          if (next_symbol_.symbol == noToken) {
            fetchNext(lexer, stack);
          }
          entry.sym.offset = next_symbol_.offset;
          entry.sym.endoffset = next_symbol_.endoffset;
        } else {
          rhs = absl::Span<const stackEntry>(&stack[0] + stack.size() - ln, ln);
          entry.sym.offset = rhs.front().sym.offset;
          entry.sym.endoffset = rhs.back().sym.endoffset;
        }
        absl::Status ret = applyRule(rule, entry, rhs, lexer);
        if (!ret.ok()) {
          return ret;
        }
        stack.resize(stack.size() - ln);
        if (debugSyntax) {
          LOG(INFO) << "reduced to: " << symbolName(entry.sym.symbol)
                    << " consuming " << ln << " symbols, range "
                    << entry.sym.offset << " to " << entry.sym.endoffset;
        }
        state = gotoState(stack.back().state, entry.sym.symbol);
        entry.state = state;
        stack.push_back(std::move(entry));
      } else if (action == -1) {
        // Shift.
        if (next_symbol_.symbol == noToken) {
          fetchNext(lexer, stack);
        }
        state = gotoState(state, next_symbol_.symbol);
        if (state >= 0) {
          stack.push_back(stackEntry{
              .sym = next_symbol_,
              .state = state,
          });
          if (debugSyntax) {
            LOG(INFO) << "shift: " << symbolName(next_symbol_.symbol) << " ("
                      << lexer.Text() << ")";
          }
          if (!pending_symbols_.empty()) {
            for (const auto& tok : pending_symbols_) {
              reportIgnoredToken(tok);
            }
            pending_symbols_.clear();
          }
          if (next_symbol_.symbol != eoiToken) {
            switch (Token(next_symbol_.symbol)) {
              case Token::JSONSTRING:
                listener_(NodeType::JsonString, next_symbol_.offset,
                          next_symbol_.endoffset);
                break;
              default:
                break;
            }
            next_symbol_.symbol = noToken;
          }
        }
      }
      if (action == -2 || state == -1) {
        break;
      }
    }

    if (state != end) {
      if (next_symbol_.symbol == noToken) {
        fetchNext(lexer, stack);
      }
      return absl::InvalidArgumentError(absl::StrFormat(
          "Syntax error: line %d: %s", lexer.Line(), lexer.Text()));
    }

    return absl::OkStatus();
  }

  template <typename Lexer>
  void fetchNext(Lexer& lexer, std::vector<stackEntry>& stack) {
    Token tok;
  restart:
    tok = lexer.Next();
    switch (tok) {
      case Token::MULTILINECOMMENT:
      case Token::INVALID_TOKEN:
        pending_symbols_.push_back(symbol{static_cast<int32_t>(tok),
                                          lexer.TokenStartLocation(),
                                          lexer.TokenEndLocation()});
        goto restart;
      default:
        break;
    }

    next_symbol_.symbol = static_cast<int32_t>(tok);
    next_symbol_.offset = lexer.TokenStartLocation();
    next_symbol_.endoffset = lexer.TokenEndLocation();
  }

  void reportIgnoredToken(symbol sym);

  template <typename Lexer>
  absl::Status applyRule(int32_t rule, stackEntry& lhs,
                         [[maybe_unused]] absl::Span<const stackEntry> rhs,
                         Lexer& lexer) {
    switch (rule) {
      case 32:
        if (AtEmptyObject(lexer, next_symbol_.symbol)) {
          lhs.sym.symbol = 23; /* lookahead_EmptyObject */

        } else {
          lhs.sym.symbol = 25; /* lookahead_notEmptyObject */
        }
        return absl::OkStatus();
      default:
        break;
    }

    if (NodeType nt = tmRuleType[rule]; nt != NodeType::NoType) {
      listener_(nt, lhs.sym.offset, lhs.sym.endoffset);
    }
    return absl::OkStatus();
  }

  symbol next_symbol_;
  // Tokens to be reported with the next shift. Only non-empty when next.symbol
  // != noToken.
  std::vector<symbol> pending_symbols_;
  absl::FunctionRef<void(NodeType, int64_t, int64_t)> listener_;

  friend std::ostream& operator<<(std::ostream& os, const Parser& parser) {
    return os << "JSONParser next " << symbolName(parser.next_symbol_.symbol)
              << " and pending " << parser.pending_symbols_.size()
              << " symbols";
  }
};

}  // namespace json
#endif  // EXAMPLES_JSON_PARSER_H_
