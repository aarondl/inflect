Fork of bitbucket.org/pkg/inflect

INSTALLATION

`go get -u github.com/aarondl/inflect`

PACKAGE

package inflect


FUNCTIONS

func AddAcronym(word string)

func AddHuman(suffix, replacement string)

func AddIrregular(singular, plural string)

func AddPlural(suffix, replacement string)

func AddSingular(suffix, replacement string)

func AddUncountable(word string)

func Asciify(word string) string

func Camelize(word string) string

func CamelizeDownFirst(word string) string

func Capitalize(word string) string

func Dasherize(word string) string

func ForeignKey(word string) string

func ForeignKeyCondensed(word string) string

func Humanize(word string) string

func Ordinalize(word string) string

func Parameterize(word string) string

func ParameterizeJoin(word, sep string) string

func Pluralize(word string) string

func Singularize(word string) string

func Tableize(word string) string

func Titleize(word string) string

func Typeify(word string) string

func Uncountables() map[string]bool

func Underscore(word string) string


TYPES

type Rule struct {
    // contains filtered or unexported fields
}
used by rulesets

type Ruleset struct {
    // contains filtered or unexported fields
}
a Ruleset is the config of pluralization rules
you can extend the rules with the Add* methods

func NewDefaultRuleset() *Ruleset
create a new ruleset and load it with the default
set of common English pluralization rules

func NewRuleset() *Ruleset
create a blank ruleset. Unless you are going to
build your own rules from scratch you probably
won't need this and can just use the defaultRuleset
via the global inflect.* methods

func (rs *Ruleset) AddAcronym(word string)
if you use acronym you may need to add them to the ruleset
to prevent Underscored words of things like "HTML" coming out
as "h_t_m_l"

func (rs *Ruleset) AddHuman(suffix, replacement string)
Human rules are applied by humanize to show more friendly
versions of words

func (rs *Ruleset) AddIrregular(singular, plural string)
Add any inconsistant pluralizing/sinularizing rules
to the set here.

func (rs *Ruleset) AddPlural(suffix, replacement string)
add a pluralization rule

func (rs *Ruleset) AddPluralExact(suffix, replacement string, exact bool)
add a pluralization rule with full string match

func (rs *Ruleset) AddSingular(suffix, replacement string)
add a singular rule

func (rs *Ruleset) AddSingularExact(suffix, replacement string, exact bool)
same as AddSingular but you can set `exact` to force
a full string match

func (rs *Ruleset) AddUncountable(word string)
add a word to this ruleset that has the same singular and plural form
for example: "rice"

func (rs *Ruleset) Asciify(word string) string
transforms latin characters like é -> e

func (rs *Ruleset) Camelize(word string) string
"dino_party" -> "DinoParty"

func (rs *Ruleset) CamelizeDownFirst(word string) string
same as Camelcase but with first letter downcased

func (rs *Ruleset) Capitalize(word string) string
uppercase first character

func (rs *Ruleset) Dasherize(word string) string
"SomeText" -> "some-text"

func (rs *Ruleset) ForeignKey(word string) string
an underscored foreign key name "Person" -> "person_id"

func (rs *Ruleset) ForeignKeyCondensed(word string) string
a foreign key (with an underscore) "Person" -> "personid"

func (rs *Ruleset) Humanize(word string) string
First letter of sentance captitilized
Uses custom friendly replacements via AddHuman()

func (rs *Ruleset) Ordinalize(str string) string
"1031" -> "1031st"

func (rs *Ruleset) Parameterize(word string) string
param safe dasherized names like "my-param"

func (rs *Ruleset) ParameterizeJoin(word, sep string) string
param safe dasherized names with custom seperator

func (rs *Ruleset) Pluralize(word string) string
returns the plural form of a singular word

func (rs *Ruleset) Singularize(word string) string
returns the singular form of a plural word

func (rs *Ruleset) Tableize(word string) string
Rails style pluralized table names: "SuperPerson" -> "super_people"

func (rs *Ruleset) Titleize(word string) string
Captitilize every word in sentance "hello there" -> "Hello There"

func (rs *Ruleset) Typeify(word string) string
"something_like_this" -> "SomethingLikeThis"

func (rs *Ruleset) Uncountables() map[string]bool

func (rs *Ruleset) Underscore(word string) string
lowercase underscore version "BigBen" -> "big_ben"


