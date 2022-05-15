package models

type FileData struct {
        Id              string `json: "id"`
        Name            string `json: "name"`
        Type            string `json: "type"`
        Hash_sum        string `json: "hash_sum"`
}

