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
%type<expr> document unordered_list_item unordered_list

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
