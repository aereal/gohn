# gohn

gohn (Hatena Notation parser written in Go) is a parser for Hatena Notation (a.k.a. はてな記法).

**:warning: This software is like a sketch. API or any design may change without notice.**

## AST definition

See [schema.json](schema.json)

## Background

Why do we implement new Hatena Notation parser although there are several implementations?

In short, all of other implementations have a problem about complexity, or (and) lack of formal specification.

### Complexity

The original Hatena Notation spec. and its implementation have some very complex features.
For example, `[URL:title]` notation that tells the parser retrieve `URL` and use retrieved page title for anchor text.

Former implementations have failed to separation of concerns.

### Lack of formal specification

As mentioned earlier, we have a specification of Hatena Notation, but it is not formal.
In other words, former implementations are also informal specification.

This means we have to know how implementations work if we want to know definition of some notations.

It also causes lost of encapsulation, so changes will be mostly breaking.

## Author

@aereal

## License

The package is available as open source under the terms of the MIT License (see LICENSE).
