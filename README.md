# omed server

the omed(open my editor) chrome plugin server side

accept request `http://127.0.0.1:8989/file?param` to open file in your editor

## we now support 

- ide: jetbrain[clion,goland,idea], vscode (make sure they are in your path)
- codebase: github(gitee), source.chromium
- operating system(only tested in macos)

## how to use

1. config your codebase in `omed-conf.yaml`  
the name of your codebase must match the url in web site
   
2. run the server in background

3. use web browser extension omed
    - right click a link 
    - right click on selected text
    
4. it will send request to the server and server try to open code in your ide
    