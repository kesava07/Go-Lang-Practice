const express = require("express");
const app = express();
const PORT = 8000;

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get('/' , (req, res) => {
    res.status(200).send("Welcome to chenna server")
})

app.get('/get', (req, res) => {
    res.status(200).send("This is GET request")
})

app.post('/post', (req, res) => {
    const reqBody = req.body
    res.status(201).send(reqBody)
})

app.post('/postForm', (req, res) => {
    res.status(201).send(JSON.stringify(req.body))
})

app.listen(PORT, ()=>{
    console.log("Server started successfully and PORT is", PORT)
})
