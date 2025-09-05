import React, { useState } from "react";

function App() {
  const [file, setFile] = useState(null);
  const [preview, setPreview] = useState(null);
  const [uploadedUrl, setUploadedUrl] = useState(null);

  const handleFileChange = (e) => {
    const f = e.target.files[0];
    setFile(f);
    if (f) setPreview(URL.createObjectURL(f));
  };

  const handleUpload = async (e) => {
    e.preventDefault();
    if (!file) return alert("choose a file");

    const formData = new FormData();
    formData.append("file", file);

    const res = await fetch("http://localhost:3030/upload", {
      method: "POST",
      body: formData,
    });

    const data = await res.json();
    if (res.ok) {
      setUploadedUrl("http://localhost:3030" + data.url);
    } else {
      alert(data.error);
    }
  };

  return (
    <div style={{ padding: 20 }}>
      <h1>Upload Demo</h1>
      <form onSubmit={handleUpload}>
        <input type="file" onChange={handleFileChange} />
        <button type="submit">Upload</button>
      </form>

      {preview && (
        <div>
          <h3>Preview before upload:</h3>
          <img src={preview} alt="preview" width="200" />
        </div>
      )}

      {uploadedUrl && (
        <div>
          <h3>Uploaded from server:</h3>
          <img src={uploadedUrl} alt="uploaded" width="200" />
        </div>
      )}
    </div>
  );
}

export default App;

