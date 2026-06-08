Syntax: 15 lines exactly, newline seperated. All shortforms used are the ones found in {INSERT LINK LATER}. Alignments are as is, eg. "tp", "ne". To define nonstandard team roles, use a | to seperate team and role, eg. "c|ww", "na|tpow". For "Or"s, use a /, eg. "tp/ts/ti" will have a 1/3 chance of any of the alignments. For Random and Common Town/Coven, use a ~ before the term. The code expects either a "r" or a "c" directly after the ~, then either "t" or "c", eg. "~cc", "~rt".

The parsing order is "|", "/", "~", then finally random role selection. Bare roles are not accepted, eg "jinx", "sc". You must specify a team if you are saying a specific role, eg "c|jinx", "na|sc".

The Any keyword is not supported.