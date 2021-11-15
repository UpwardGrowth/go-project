﻿# go-project
 
### How to request

### 1.Get all items

curl http://localhost:8000/albums

### 2.Add a new item

curl http://localhost:8000/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
    
### 3.Get a specific item

curl http://localhost:8000/albums/2
