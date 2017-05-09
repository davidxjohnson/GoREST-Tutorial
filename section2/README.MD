# GoREST-Tutorial
## Section 2: Setup Your Environment
**Note:**  This tutorial leans heavily toward a Linux environment and thus the instruction and examples are Linux oriented. All of the tools listed below will install on Windows, but the author makes no allowance for potential differences between Windows and Linux environments. A Windows developer should have no difficulty translating where necessary.

### A git client
Installing a git client is not absolutely necessary but highly recommended. The reader is encouraged to clone or fork the repository. At a minimum, git will save you a lot of cut/paste.

Find the git client for your environment [here](https://git-scm.com/downloads). There are numerous GUI interfaces for git, although the reader is encouraged to take the plunge and learn the git command line [basics](https://git-scm.com/doc).

### Go compiler and tools
Setting up the go tools is essential for [section3](../section3) and beyond. You can get started [here](https://golang.org/doc/install). Once you have your go environment setup, it's recommended that you make a `src` directory under `$GOPATH` and `cd` to that directory before cloning this repo.

```bash
mkdir $GOPATH/src
cd $GOPATH/src
git clone https://github.com/davidxjohnson/GoREST-Tutorial.git
cd GoREST-Tutorial
```

### Atom editor and the API Workbench
There are a number of great editors available and everyone has their favorite, but the main reason the [Atom editor](https://atom.io/) is recommended for this tutorial is that it supports the [API Workbench](http://apiworkbench.com/), a [RAML](http://raml.org/) editor add-on to Atom. The reader will need these to complete [section 4](../section4). There are also some [good add-ons](https://atom.io/packages/search?utf8=%E2%9C%93&q=golang) for Go.

### NodeJS with raml2html and RAML Markdown theme
Once we have our API design, we'll want to publish it to HTML and or markdown formats so that others can quickly understand how to use our new API. Some of the best tools for this are written in NodeJS, and it's not too difficult to install and use them.

First is to [install NodeJS](https://nodejs.org/en/) to get the npm utility installed. NPM is a module installer for the NodeJS framework. We'll use NPM to install [raml2html](https://www.npmjs.com/package/raml2html). Be sure to install the RAML Markdown theme as well.

### REST Testing Tools
Finally, we'd like (but don't absolutely need) some testing tools that are a little more elegant than curl or wget. There are two in particular: We can use [Postman](https://chrome.google.com/webstore/detail/postman/fhbjgbiflinjbdggehcddcbncdddomop?hl=en) for Chrome browsers (preferred) or [REST Client](https://addons.mozilla.org/en-US/firefox/addon/restclient/) for FireFox.

### What's next?
Now that we have the basic tools installed, let's move on to [section 3](../section3) to learn some REST programming concepts.
