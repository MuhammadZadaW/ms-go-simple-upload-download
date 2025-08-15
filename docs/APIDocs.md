# API Documentation — ms-go-simple-upload-download

## Upload File

**Endpoint:**

```
POST /upload
```

**Description:**
Uploads a single file with an optional custom name.

**Request:**
- **Content-Type:** `multipart/form-data`
- **Form Fields:**
  - `file` (file, required) — The file to upload.
  - `name` (string, optional) — Custom filename without extension.

**Response:**
- **200 OK**
```json
{
     "status": 200,
     "message": "File uploaded successfully"
}
```
- **400 OK**
```json
{
     "status": 400,
     "message": "File is required"
}
```

**Endpoint:**

```
POST /multi-upload
```

**Description:**
Uploads multiple files in a single request, returns status for each file.

**Request:**
- **Content-Type:** `multipart/form-data`
- **Form Fields:**
  - `files` (array, required) — files to upload.

**Response:**
- **207 OK**
```json
{
     "status": "partial",
     "results": [
          {
               "originalName": "file.txt",
               "status": "success",
               "message": "File uploaded successfully"
          },
          {
               "originalName": "bigfile.zip",
               "status": "error",
               "message": "File size is too large"
          }
     ]
}
```

## Download File

**Endpoint:**

```
GET /download/{filename}
```

**Description:**
Downloads a file by its name.

**Response:**
- Returns the file contents if found.
- **404 NOT FOUND**
```json
{
     "status": 404,
     "message": "File not found"
}
```

**Endpoint:**

```
POST /list
```

**Description:**
Retrieves a list of all files currently stored on the server.

**Response:**
- **200 OK**
```json
{
     "status": 200,
     "message": "Files listed successfully",
     "data": [
          "download.txt"
     ]
}
```