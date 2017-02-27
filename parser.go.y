%{
package main

%}

%union{
  token Token
  block Block
  blocks []Block
  inline Inline
  inlines []Inline
}

%token<token> TEXT
%token UNORDERED_LIST_MARKER CR LBRACKET RBRACKET
%type<block> block unordered_list_item unordered_list line
%type<blocks> blocks
%type<inline> inline inline_text inline_http
%type<inlines> inlines

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
    empty_line
    {
      $$ = Line{inlines: []Inline{}}
    }
    | inlines CR
    {
      $$ = Line{inlines: $1}
    }

empty_line:
          CR

inlines:
       inline
       {
        $$ = []Inline{$1}
       }
       | inline inlines
       {
        $$ = append([]Inline{$1}, $2...)
       }

inline:
      inline_http
      | inline_text

inline_text:
      TEXT
      {
        $$ = InlineText{literal: $1.literal}
      }

inline_http:
           LBRACKET TEXT RBRACKET
           {
            $$ = InlineHttp{url: $2.literal}
           }

unordered_list:
              unordered_list_item
              {
                $$ = UnorderedList{Items: []UnorderedListItem{$1.(UnorderedListItem)}}
              }
              | unordered_list_item unordered_list
              {
                items := $2.(UnorderedList).Items
                list := UnorderedList{Items: append([]UnorderedListItem{$1.(UnorderedListItem)}, items...)}
                $$ = list
              }

unordered_list_item:
                   UNORDERED_LIST_MARKER TEXT CR
                   {
                    $$ = UnorderedListItem{Text: $2.literal}
                   }

%%
