# Templ-gen: a Templ/GO generator with Tailwind CSS and TypeScript support

Templ-gen requires these CLI tools:
- air
- Tailwind CSS
- Templ
- TSC
<!--
To install CLI:
```bash
mkdir ~/.templ-gen && mkdir ~/.templ-gen/bin
curl -sSfL https://raw.githubusercontent.com/bradscottwhite/templ-gen/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```
-->

To run:

```bash
mkdir templ-demo && cd templ-demo

templ-gen i

air

# Generate new component
templ-gen c newComponent

# Generate new page
templ-gen p about
```
