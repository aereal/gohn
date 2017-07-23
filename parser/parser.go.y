%{
package parser

%}

%union{
  token Token
  block Block
  blocks []Block
  inline Inline
  inlines []Inline
  url string
  http_option string
  http_options HttpOptions
  reference Reference
  depth int
}

%token<token> TEXT
%token UNORDERED_LIST_MARKER ORDERED_LIST_MARKER CR LBRACKET RBRACKET LT GT HEADING_MARKER COLON
%type<block> block unordered_list_item unordered_list ordered_list ordered_list_item line quotation heading
%type<blocks> blocks
%type<inline> inline inline_text inline_http
%type<inlines> inlines
%type<url> url
%type<http_options> http_options
%type<http_option> http_option
%type<reference> reference quotation_prefix
%type<depth> unordered_list_markers
%type<depth> heading_prefix

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
        | ordered_list
        {
          $$ = $1
        }
        | quotation
        {
          $$ = $1
        }
        | heading
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
      $$ = Line{Inlines: []Inline{}}
    }
    | inlines CR
    {
      $$ = Line{Inlines: $1}
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
        $$ = InlineText{Literal: $1.literal}
      }

inline_http:
           LBRACKET reference RBRACKET
           {
            $$ = InlineHttp{Reference: $2}
           }

reference:
  url
  {
    $$ = Reference{Url: $1}
  }
  | url http_options
  {
    $$ = Reference{Url: $1, Options: $2}
  }

url: TEXT
   {
    $$ = $1.literal
   }

http_options:
  http_option
  {
    $$ = []string{$1}
  }
  | http_option http_options
  {
    options := $2
    $$ = append([]string{$1}, options...)
  }

http_option: COLON TEXT
  {
    $$ = $2.literal
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
                   unordered_list_markers inlines CR
                   {
                    $$ = UnorderedListItem{Depth: $1, Inlines: $2}
                   }

unordered_list_markers:
                     unordered_list_marker
                     {
                      $$ = 1
                     }
                     | unordered_list_marker unordered_list_markers
                     {
                      $$ = $2 + 1
                     }


unordered_list_marker:
                     UNORDERED_LIST_MARKER

ordered_list:
              ordered_list_item
              {
                $$ = OrderedList{Items: []OrderedListItem{$1.(OrderedListItem)}}
              }
              | ordered_list_item ordered_list
              {
                items := $2.(OrderedList).Items
                list := OrderedList{Items: append([]OrderedListItem{$1.(OrderedListItem)}, items...)}
                $$ = list
              }

ordered_list_item:
                   ORDERED_LIST_MARKER inlines CR
                   {
                    $$ = OrderedListItem{Inlines: $2}
                   }

quotation:
         quotation_prefix blocks quotation_suffix
         {
          $$ = Quotation{Cite: $1, Content: $2}
         }

quotation_prefix:
                GT reference GT CR
                {
                  $$ = $2
                }
                | GT GT CR
                {
                  $$ = Reference{}
                }

quotation_suffix:
                LT LT CR

heading:
       heading_prefix inlines CR
       {
        $$ = Heading{Level: $1, Content: $2}
       }

heading_prefix:
              HEADING_MARKER
              {
                $$ = 1
              }
              | HEADING_MARKER heading_prefix
              {
                $$ = $2 + 1
              }

%%
