%{
package main

%}

%union{
  token Token
  expr Expr
  exprs []Expr
}

%token<token> TEXT
%token UNORDERED_LIST_MARKER
%type<expr> block unordered_list_item unordered_list
%type<exprs> blocks

%%

blocks:
      block
      {
        yylex.(*Lexer).result = []Expr{$1}
      }

block:
        unordered_list
        {
          $$ = $1
        }

unordered_list:
              unordered_list_item
              {
                $$ = UnorderedListExpr{items: []UnorderedListItemExpr{$1.(UnorderedListItemExpr)}}
              }
              | unordered_list_item unordered_list
              {
                items := $2.(UnorderedListExpr).items
                list := UnorderedListExpr{items: append([]UnorderedListItemExpr{$1.(UnorderedListItemExpr)}, items...)}
                $$ = list
              }

unordered_list_item:
                   UNORDERED_LIST_MARKER TEXT
                   {
                    $$ = UnorderedListItemExpr{text: $2.literal}
                   }

%%
