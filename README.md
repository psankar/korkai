#korkai கொற்கை - A Tamil corpus builder
This project helps in building a corpus of Tamil words. You can use the programs
to analyze wordpress, blogger and wikipedia archives and generate two kinds of
files. Both the files will contain the list of unique Tamil words used in the
archives where one is sorted alphabetically and the other is sorted based on
the usage count of the words.

#License
All the source files are licensed under Creative Commons Zero License
This is also called public domain in some countries
For more information see:	http://creativecommons.org/publicdomain/zero/1.0/
For full license text see:	http://creativecommons.org/publicdomain/zero/1.0/legalcode

Submit your patches, changes or pull-requests only if you are ready to give them
under the Creative Commons Zero License.


#Installation
Currently you can install korkai by only building from the source
TODO: Give the location of the binaries and the installation instructions

#Installation from the sources on Linux/Mac OS machines
* Install the Go Programming language from http://golang.org
* Download your blogger/wordpress/wikipedia archive files
* $> go run wikipedia-archives.go <wikipedia-archive>.xml
* $> go run blogger-xmldump-parser.go <blogger-archive>.xml
* $> go run wordpress-xmldump-parser.go <wordpress-archive>.xml
* $> ls *.tokens *.stats

The *.tokens files contain the alphabetically sorted unique Tamil words
The *.stats files contain the usage-count-sorted unique Tamil words

If you have multiple blog archives of the same type (blogger or wordpress) you
can pass all the filenames at the same time as parameters.

We use the words in the blogposts only and NOT the comments for our analysis

#Downloading Archives
###blogger archives
Goto blogger.com -> Settings -> Other -> Export blog
The saved archive will mostly be like blog-MM-DD-YYYY.xml
###wordpress archives
TODO: Get the instructions from some wordpress user and fill here
###wikipedia archives
Download http://dumps.wikimedia.org/tawiki/latest/tawiki-latest-pages-articles.xml.bz2
Unzip the above compressed archive to extract the .xml file
Linux users can unzip the archive file using tar or file-roller
Windows users can unzip the archive by http://gnuwin32.sourceforge.net/packages/bzip2.htm

#The Name Korkai
The name Korkai refers to http://en.wikipedia.org/wiki/Korkai

Korkai was the name of a once-coastal town, a natural harbor that was once a 
pearl fishery. It was the source of the world's biggest pearls once. Romans, 
Chinese and every major civilization (roughly around 3rd BC) were eager to come
to Korkai to get the beautiful pearls.

For more information about Korkai, read the novel Raja Muthirai by the renowned
Tamil author Chaandilyan/Sandilyan https://www.goodreads.com/book/show/7715528-raja-muththirai

Just as how the Korkai helped in getting the best pearls of then world, my tool
korkai aims to identify the pearls, Tamil Words from documents.

#Credits
The following people helped me test the program by providing me with their blog
archives. A lot of thanks to them.

* Karki Bava - http://iamkarki.blogspot.com
* Chokkan - http://nchokkan.wordpress.com
* Priya Kathiravan - http://priyakathiravan.blogspot.com
* Elavasa Kothanar - http://elavasam.blogspot.com
* Ragavan - http://gragavan.blogspot.com
* Saravanakarthikeyan - http://www.writercsk.com

Thanks to Sonia Keys and Kyle Lemmons from the golang mailing lists for helping 
me with various nittygritties of the Go programming language.

#Developer Notes
A hashtable is used to store the words in memory before dumping the data to the
files. We can improve the performance substantially by using a Trie instead of a
hastable. However, we have used a hashtable to keep the code short and more
easily understandable for anyone with just the programming language knowledge. 

Using the built in datastructures makes the maintenance of the code easier too.
However, if you are planning to use this in a production machine that requires 
realtime responses, you are recommended to change the hashtable to something 
that will be more fitting for your needs.

The wikipedia extractor program will use goroutines and hence all the CPUs in
your machine, thereby increasing the performance. In my thinkpad T430 with 4
CPUs and 4 GB RAM, it finishes the word extraction part in about 50 seconds.

The extractor programs for blogger and wordpress are not parallelized since the
load is usually very less and the benefits of parallellizing may not be worth 
the effort. This too may be done sometime in the future though, depending on the
need and more importantly how bored I am.

#TODO
* Translate error strings and messages to Tamil
* Make installable binaries for non-technical users who cannot build from source
* Identify way to merge stats of two or more archives intelligently
