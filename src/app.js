const express = require('express')
const multer = require('multer')
const upload = multer({dest: 'uploadeds/'})

const app = express()

app.get('/', function (req, res) {
    res.send('Hello World!')
})

app.post('/*', upload.any(), function (req, res) {
    console.log(req.files[0].originalname)
    res.send(req.files[0].fieldname)
})

    

app.listen(3000, function () {
      console.log('Example app listening on port 3000!')
})
