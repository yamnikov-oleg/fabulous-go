# Building Faboulous Go

To simply rebuild the list, clone the repo and run `make` in the root of it. Note that:

* It will prompt for your **GitHub credentials**. They are used for retrieving repositories info with higher rate.
* It will take time, around **20 minutes** or more.

Now, details.

## raider

This tool is used to grab the [awesome-go](https://github.com/avelino/awesome-go) list, retrieve info about each repo and save it to a json file.

To run the raider only, use `make raid` command. This will create `tmp/repos.json` file, which is used in the next step.

## exhibitor

This tool parses `repos.json` and displays it to a markdown file. The template for exhibitor is placed into `exhibitor/TEMPLATE.md`.

To run the exhibitor only, type `make exhibit`. Note, that it will require `tmp/repos.json` to exist.

## cleaning

If the toolchain crashes, run `make cleanup` to remove all the temporary files used in grabbing.

Note that `repos.json` persist between runs, because it's very time-consuming to restart raider every time, so sometimes it's easier to simply reprocess the grabbing results into a better view without regrabbing them.
