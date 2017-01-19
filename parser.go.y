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
%type<exprs> unordered_list
%type<expr> document unordered_list_item

%%

document:
        unordered_list
        {
          $$ = $1
          yylex.(*Lexer).result = $$
        }

unordered_list:
              unordered_list_item
              {
                $$ = []Expr{$1}
              }
              | unordered_list_item unordered_list
              {
                $$ = append([]Expr{$1}, $2...)
              }

unordered_list_item:
                   UNORDERED_LIST_MARKER TEXT
                   {
                    $$ = UnorderedListItemExpr{text: $2.literal}
                   }

%%
