# srt-order

Quick program to fix out of order nichijou subtitles I downloaded.

The program simply parses each `Block`, then sorts by `start` time and writes to stdout / to a file, setting the appropriate sequence number.

This is no by any means a decent parser. Don't use it.

# acknowledgements

big thanks to [konifar/go-srt](https://github.com/konifar/go-srt/tree/master), which saved me some time from writing some parsing code for the timestamp fields.
