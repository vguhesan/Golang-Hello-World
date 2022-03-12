# Golang-Hello-World
Provides a quick starter template for a Golang Hello World project. 

# Why?
If you're looking to test out some simple logic and you want to do that locally in a hello-world Golang project then this will allow you to quickly bring down this repo and test out your code and you can dispose it or store it in your own custom repo.

# Why another hello-world?
Most hello-world implementations gives you a main.go that prints a "Hello World". I wanted something with a bit more organization and package structure as a starting point. I would highly recommend you take a look at this Golang Standards Layout to get a understanding of recommended ways yo organize your Golang projects:

[link](https://github.com/golang-standards/project-layout)

# Getting Started

You can simply clone this repo down using something like `git clone {repo-name}`. This will download the entire repository with all branches, git commits and history which you don't really need for trying out something locally. This is where "[degit](https://github.com/Rich-Harris/degit)" comes in. "degit" makes copies of git repositories. By the way, "degit" is created by Rich Harris that created one of my favorite frameworks, namely "[Svelte](https://svelte.dev/)". If you don't know what it is and you're into Angular or React programming you should check it out. 

## Prerequisite
Make sure your have NodeJS installed. This gives you npx (npx executes <command> either from a local node_modules/.bin, or from a central cache, installing any packages needed in order for <command> to run. With npm you are downloading the package and all it's dependencies. With npx - you don't download it. You simply pull and execute it.)

```
# Ensure you have npx (see prerequisite section above)
# To verify you have npx, run "which npx" or "npx --version"
npx degit https://github.com/vguhesan/Golang-Hello-World my-golang-hello-world
cd my-golang-hello-world
# Edit go.mod line 0 "example/hello" to something like "github.com/{username}/myapp"
go build
# or 
go run main.go
```


