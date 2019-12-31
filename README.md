# gorename
A bulk renaming utility written in golang for the command line

## Demonstration
<div style="width:100%;height:0px;position:relative;padding-bottom:56.388%;"><iframe src="https://streamable.com/s/15m6e/kwujnt" frameborder="0" width="100%" height="100%" allowfullscreen style="width:100%;height:100%;position:absolute;left:0px;top:0px;overflow:hidden;"></iframe></div>

[Here's the streamable link](https://streamable.com/15m6e)

## -h output
```
usage: gorename [-h|--help] -e|--command (rename|prep) [-c|--contentDir
                "<value>"]

                Bulk renaming utility written in golang

Arguments:

  -h  --help        Print help information
  -e  --command     The command you want to run. `rename` uses a buffer file to
                    rename files. `prep` creates a buffer file
  -c  --contentDir  Directory of files to be renamed

```
## Example
Create a buffer file to edit
```
gorename -e prep -c path/to/dir/containing/to-be-renamed-files
```
OR if you're on a unix based machine, a simpler command would be...
```
ls path/to/dir/containing/ >> buffer.txt
```
This creates a `buffer.txt` file with the names of all the files in the given directory in your current working directory. 

#### NOTE: run this utility OUTSIDE of the directory with all the files you want to rename

#### Then edit the filenames in the buffer file to what you want your directory to look like (order will be preserved)

Use buffer file to rename 
```
gorename -e rename -c path/to/dir/containing/to-be-renamed-files
```
and gorename will remove the buffer file for you

I AM NOT RESPONSIBLE FOR ANY LOSS OF DATA USE AT YOUR OWN CAUTION, READ MIT LICENSE