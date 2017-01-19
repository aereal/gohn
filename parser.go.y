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
%type<block> block unordered_list_item unordered_list line
%type<blocks> blocks

%%

blocks:
      block
      {
        $$ = []Block{$1}
        yylex.(*Lexer).result = $$
      }
      | block blocks
      {
        $$ = append([]Block{$1}, $2...)
        yylex.(*Lexer).result = $$
      }

block:
        unordered_list
        {
          $$ = $1
        }
        | line
        {
          $$ = $1
        }

line:
    TEXT
    {
      $$ = Line{text: $1.literal}
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
