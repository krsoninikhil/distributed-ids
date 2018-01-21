var express = require('express');
var app = express();

app.use(express.static('public'));
app.get('/index.html', function (req, res) {
   res.sendFile( __dirname + "/" + "index.html" );
})


const { exec}  = require('child_process');
app.get('/process_get', function (req, res) {
   // Prepare output in JSON format
   //
    const execSync = require('child_process').execSync;
        code = execSync('node ../fabcar/invoke.js admin ' +req.query.uin+ ' "' + req.query.name+ '" ' + req.query.age+' "'+ req.query.dob+'" ' + req.query.contact)

   exec('node ../fabcar/invoke.js admin user1 req.query.name req.query.age req.query.dob req.query.contact')
      response = {
      uin:req.query.uin,
      name:req.query.name,
      age:req.query.age,
      date_of_birth:req.query.dob,
      contact:req.query.contact
   };
   console.log(response);
   res.end(JSON.stringify(response));
})

app.get('/signup', function (req, res) {
   // Prepare output in JSON format
   const execSync = require('child_process').execSync;
	code = execSync('node ../fabcar/registerUser.js ' + req.query.inputname)
      response = {
      name:req.query.inputname
   };
   console.log(response);
   res.end(JSON.stringify(response));
})



var server = app.listen(8081, function () {
   var host = server.address().address
   var port = server.address().port
   console.log("Example app listening at http://%s:%s", host, port)

})
