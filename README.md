# Mindctrl Web Extension

This is a simple web extension for controlling a web browser by external
programs and scripts.

The tool is used personally to script the browser for web scraping and
automating some complex workflow on websites.

## Rationale

There are a number of libraries that involves browser automation. But
they are not entirely fit for my purpose.

First of all, most of them uses either Chrome remote debugging protocol
and/or Firefox marionette to control the browser. However, when they are
enabled, the browser will set the navigator.webdriver and reveal to the
remote sites that they are being programmatically controlled. Some
websites, like those protected y Cloudflare, will reject such browsers.

Moreover, many of the libraries are designed for running their own browser
instances, instead of working with an existing browser instance.

## Status

At the moment the project is more or less ready for personal use. However,
it should be considered alpha quality. Expect rough edges, bugs and lack
of backward/foreward compatibility. **DO NOT depend on the project for
anything mission critical.**

The project is mainly tested against Firefox. There are some testing
against Chromium. Other browsers are probably not supported.

## Security

As the extension enables external program to control a browser, it can
be abused to do anything the browser can do. **DO NOT use, or even
install this extension in browsers that is used for sensitive tasks
such as online banking.**

## Dependencies

The extension is expected to be built on a UNIX environment. The extension
requires *make*, *zip*, *node.js* and *npm* to build; and for the client
*golang 1.18* is required.

## Build

The extension can be built by running make. The targets firefox and chrome
build the extension for each browser respectively. The target client
builds the mindctrl command line tool. All output can be found in the
dist subdirectory.

Note that the extension are not packed.

## Installation

The extension depends on an intermediate MQTT server for communcation
with the clients. Furthermore, the server must support websocket
connection.

The extension is for personal use and therefore it is not uploaded to
any extension store. Instead it can be installed temporarily by
the procedure [here](https://extensionworkshop.com/documentation/develop/temporary-installation-in-firefox/)
(for Firefox) and [here](https://developer.chrome.com/docs/extensions/mv3/getstarted/development-basics/#load-unpacked) (for Chrome).

## Others

The project is mostly developed for personal use. Do not expect
active development or project management at all. In particular,
Github issue will be ignored.

