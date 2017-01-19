%{
package main

%}

%union{
  token Token
  block Block
  blocks []Block
}

%token<token> TEXT
%token UNORDERED_LIST_MARKER
%type<block> block unordered_list_item unordered_list
%type<blocks> blocks

%%

blocks:
      block
      {
        yylex.(*Lexer).result = []Block{$1}
      }

block:
        unordered_list
        {
          $$ = $1
        }

unordered_list:
              unordered_list_item
              {
                $$ = UnorderedList{items: []UnorderedListItem{$1.(UnorderedListItem)}}
              }
              | unordered_list_item unordered_list
              {
                items := $2.(UnorderedList).items
                list := UnorderedList{items: append([]UnorderedListItem{$1.(UnorderedListItem)}, items...)}
                $$ = list
              }

unordered_list_item:
                   UNORDERED_LIST_MARKER TEXT
                   {
                    $$ = UnorderedListItem{text: $2.literal}
                   }

%%
