# evolve

An evolution engine loosely based on Biomorphs from Dawkins' *The Blind
Watchmaker*.

The core type is `Evolution`, which a pointer to an instance of can be obtained
via the `evolve.New(start *Genome, development func(*Genome) *Individual)`
function. Once an instance `e` is obtained, `e.Evolve(size int)` can be called
to retrieve the next generation of `Individual`s, and then the generation's
survivor can be passed to `e.Select(survivor *Individual)`.

The `development` function passed to `New` is in charge of taking a `Genome` and
developing it into an `Individual`, which includes both a genotype and a
phenotype. The phenotype is declared as `interface{}`, but would most likely be
some sort of image (as it was for Biomorphs).  It is not used by `Evolution` in
any way; it's just for the natural selection process (selecting the `Individual`
returned from `Evolve` to pass to `Select`) that happens via some external
process.

Behind the scenes there is also a `Reproduce` function which takes a `Genome`
and creates a child `Genome` from it with exactly one random point mutation.

To complete the Biomorphs implementation, you'd need to implement `development`
to draw the tree-ish structures, and then create some sort of user interface
that would present the tree-ish images to the user to let them select.

The latest generation's surviving `Individual` is always available at `e.Latest`
and it's complete ancestry is in `e.Ancestry`.  For space reasons, only the
`Genome`s are kept in the ancestry, but they can be developed into `Individual`s
if needed by using the development function at `e.Development`.
