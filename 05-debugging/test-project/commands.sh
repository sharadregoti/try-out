docker build -t sharadregoti/go-debug:latest
docker run --name go-debug -p 4040:4040 -p 8090:8090 sharadregoti/go-debug:latest
kubectl port-forward go-debug-container-747bdd978c-hltqw 4040:4040 8090:8090

# {
#     // Use IntelliSense to learn about possible attributes.
#     // Hover to view descriptions of existing attributes.
#     // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
#     "version": "0.2.0",
#     "configurations": [
#         {
#             "name": "Connect to server",
#             "type": "go",
#             "request": "attach",
#             "mode": "remote",
#             "port": 4040,
#             "host": "127.0.0.1"
#         },

#         {
#             "name": "Launch Package",
#             "type": "go",
#             "request": "launch",
#             "mode": "auto",
#             "program": "${workspaceFolder}"
#         }
#     ]
# }