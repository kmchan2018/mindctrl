

icons := icons/icon-normal.png icons/icon-dark.png
pages := frontend.html options.html
stylesheets := common.css
javascripts := background.js frontend.js options.js content_scripts/query.js content_scripts/wait.js
typescripts := $(shell cd extension && find . -name '*.ts')
gofiles := $(shell cd client && find . -name '*.go')
artifacts := manifest.json $(icons) $(pages) $(stylesheets) $(javascripts)


.PHONY: all firefox chrome client check
.SECONDARY: $(addprefix build/,$(stylesheets))


all: firefox chrome client

client: dist/mindctrl

chrome: $(addprefix dist/chrome/,$(artifacts))

firefox: dist/firefox.zip

dist/mindctrl: $(addprefix client/,$(gofiles)) client/go.mod client/go.sum
	cd client && go build -o ../dist/mindctrl cmd/mindctrl.go

dist/firefox.zip: $(addprefix dist/firefox/,$(artifacts))
	cd dist/firefox && zip -r -FS ../firefox.zip $(artifacts)

dist/firefox/manifest.json: extension/manifest.firefox.json
	install -m 0755 -d $(dir $@)
	install -m 0644 $< $@

dist/firefox/%.css: build/%.css
	install -m 0755 -d $(dir $@)
	install -m 0644 $< $@

dist/firefox/%.js: build/%.js
	install -m 0755 -d $(dir $@)
	install -m 0644 $< $@

dist/firefox/%: extension/%
	install -m 0755 -d $(dir $@)
	install -m 0644 $< $@

dist/chrome/manifest.json: extension/manifest.chrome.json
	install -m 0755 -d $(dir $@)
	install -m 0644 $< $@

dist/chrome/%.css: build/%.css
	install -m 0755 -d $(dir $@)
	install -m 0644 $< $@

dist/chrome/%.js: build/%.js
	install -m 0755 -d $(dir $@)
	install -m 0644 $< $@

dist/chrome/%: extension/%
	install -m 0755 -d $(dir $@)
	install -m 0644 $< $@

build/%.css: extension/%.tw.css $(addprefix extension/,$(pages)) tailwind.config.js node_modules/.bin/tailwindcss
	install -m 0755 -d build
	./node_modules/.bin/tailwindcss -i $< -o $@

$(addprefix build/,$(javascripts)) &: $(addprefix extension/,$(typescripts)) webpack.config.js tsconfig.json node_modules/.bin/tsc node_modules/.bin/webpack
	install -m 0755 -d build
	./node_modules/.bin/tsc --noEmit
	./node_modules/.bin/webpack

node_modules node_modules/.bin/tailwindcss node_modules/.bin/tsc node_modules/.bin/webpack &:
	npm install


check:
	node_modules/.bin/tsc --noEmit


