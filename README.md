# krex

Kubernetes Resource Explorer

Krex works by building a directional graph (digraph) in memory of various Kubernetes resources, and then giving you the graph one layer at a time to explore using an interactive drop down menu. Explore Kubernetes right from your terminal.

## Get Involved!

Join us in the `#krex` channel in the [Kubernetes Slack community](http://slack.k8s.io/).

## Current state of krex

Handy tool for exploring applications in Kubernetes

## Future of krex

Global tool for exploring all things in Kubernetes

# Building krex

### Mac OSX

We use the C `ncurses` tool internally in Krex to navigate the terminal.

First install `ncurses`

```bash
brew install ncurses
```

Export the `PKG_CONFIG_PATH` variable

```bash
export PKG_CONFIG_PATH=/usr/local/opt/ncurses/lib/pkgconfig
```

Then use symbolic links to adjust the library for `pkg-config`. (Note: more information can be found [here]https://gist.github.com/cnruby/960344).

```bash
ln -s /usr/local/opt/ncurses/lib/pkgconfig/formw.pc /usr/local/opt/ncurses/lib/pkgconfig/form.pc
ln -s /usr/local/opt/ncurses/lib/pkgconfig/menuw.pc /usr/local/opt/ncurses/lib/pkgconfig/menu.pc
ln -s /usr/local/opt/ncurses/lib/pkgconfig/panelw.pc /usr/local/opt/ncurses/lib/pkgconfig/panel.pc
```

Then build the binary

```bash
make build
```

and add to your path

```bash
mv krex /usr/local/bin
```