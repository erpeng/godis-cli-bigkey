# godis-cli-bigkey
find big keys in redis using rdb file
# usage
 copy rdb file to project root,then execute
       
       go run godis-cli-bigkey.go
       
 you can use -h to find other options
        
       go run godis-cli-bigkey.go -h

        -debug
        open debug mode
        -topn int
        output topn keys (default 100)
        -totallen
        get total len (key and meta) or only value len (default true)
