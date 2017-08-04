# GoREST-Tutorial
## Section 1 - Introduction
### The Goals of This Tutorial

#### This tutorial is intended to:
* Promote the use of Go for systems programming (particularly REST API and micro-services that are used to extend and automate back-end infrastructure).
* Reeducate developers about the foundations of web tech. It's not uncommon for programmers know how to use web tools while at the same time having a limited understanding of how the web actually works.
* Promote REST programming best practices as they pertain to Go. REST standards are not codified, but there is a growing consensus.
* Promote a top-down approach to REST micro-services. The "let's just play with Go" bottom-up approach frequently results in little or no documentation and numerous rewrites of code.

#### This tutorial is _not_:
* A general tutorial on Go. There are numerous resources (some listed at the end of this section) that are very adequate for that purpose.
* A comprehensive guide to all things REST. We cover only the the most common scenarios to help get the reader up and going with REST and Go.

### Why Go?
The key to understanding what Go is generally good or not good for depends on an understanding of what problems Go is designed to solve. Perhaps the best way to understand Go's heritage is to [read about it](https://talks.golang.org/2012/splash.article) from one of the Go language designers at Google, [Rob Pike](https://research.google.com/pubs/r.html). It's an interesting and eye-opening case study of how Google solved some of its large-scale software development problems by using Go.

What makes Go particularly appealing for REST micro-services is that it makes writing them relatively easy, the programs compile to a relatively small binary (making it particularly good for a Docker container or Raspberry PI), it performs well and it's fairly OS and hardware agnostic.

Even if convinced that Go were a compelling systems development language, nobody wants to invest time into a fad language with little support or momentum. There's little doubt Go is [growing in popularity](http://www.infoworld.com/article/3130072/application-development/with-help-from-docker-googles-go-jumps-in-popularity.html) but what removes Go from the "fad" category is the wide-scale adoption by key open-source players and large corporations. The Go community is growing along with numerous compelling open-source technologies written in Go.

Here are a few examples:
* [Kubernetes](https://github.com/kubernetes/kubernetes) - the core of Google's compute engine and Docker container orchestrator which has [gone mainstream](https://kubernetes.io/case-studies/). The GitHub project has been forked over 8000 times and is the basis for container cluster management solutions in the OpenShift and CoreOS projects.
* [Docker](https://github.com/docker) - a game-changing container technology with numerous projects written in Go.
* [OpenShift Origin](https://github.com/openshift/) - a single pane of glass for Docker deploy, scale, monitoring, logging etc. written in Go. This project is the open-source upstream to RedHat's Openshift Enterprise.
* CoreOS has [numerous](https://github.com/coreos/) widely-used infrastructure projects written in Go.
* HashiCorp has [several popular](https://github.com/hashicorp) infrastructure projects written in Go.
* There are [many, many more](https://github.com/avelino/awesome-go).
* Amazon supports a [Go SQK](https://aws.amazon.com/sdk-for-go/) for the AWS API.

### Why micro-services?
Ok, just because writing micro-services is easy with Go doesn't mean that one should go hog-wild with micro-services. They have their place with distinct advantages and disadvantages.

### Go minimum essentials:
