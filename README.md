# gorename
A bulk renaming utility written in golang for the command line

## Example
Create a buffer file to edit
```
gorename -e prep -c path/to/dir/containing/to-be-renamed-files
```
This creates a `buffer.txt` file in your current working directory. 

#### NOTE: run this utility OUTSIDE of the directory with all the files you want to rename

#### Then edit the filenames in the buffer file to what you want your directory to look like (order will be preservered)

Use buffer file to rename 
```
gorename -e rename -b buffer.txt -c path/to/dir/containing/to-be-renamed-files
```

delete the buffer.txt file and you're done

I AM NOT RESPONSIBLE FOR ANY LOSS OF DATA USE AT YOUR OWN CAUTION, USING A BLANK `buffer.txt` FILE MAY BREAK YOUR FILES