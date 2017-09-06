<?php
$uploadOk = 1;
$fileExtension = pathinfo(basename($_FILES["uploadedFile"]["name"]),PATHINFO_EXTENSION);
// Check file size
// if ($_FILES["uploadedFile"]["size"] > 500000) {
//     echo "Sorry, your file is too large.";
//     $uploadOk = 0;
// }

$date = new DateTime();
$target_dir = "uploaded_file/";
$target_file = $target_dir . "file-" . $date->getTimestamp() . "." . $fileExtension;

if(!file_exists($target_dir)) {
    mkdir($target_dir);
}
while (file_exists($target_file)) {
    $target_file = $target_dir . "file-" . $date->getTimestamp() . "." . $fileExtension;
}
if (move_uploaded_file($_FILES["uploadedFile"]["tmp_name"], $target_file)) {
    $target_type = explode("/", mime_content_type($target_file), 2)[0];
    matlab_analysis($target_file, $target_type);
} else {
    echo "Sorry, there was an error uploading your file.";
}


function matlab_analysis($filepath, $filetype) {
    $my_address = "127.0.0.1";
    $matlab_port = 3000;
    $matlab_address = "127.0.0.1";    

    /* Create a TCP/IP socket. */
    $socket = socket_create(AF_INET, SOCK_STREAM, SOL_TCP);
    if ($socket === false) {
        echo "socket_create() failed: reason: " . socket_strerror(socket_last_error()) . "\n";
    }
    
    $result = socket_connect($socket, $matlab_address, $matlab_port);
    if ($result === false) {
        echo "socket_connect() failed.\nReason: ($result) " . socket_strerror(socket_last_error($socket)) . "\n";
    }
    
    sleep(1);
    $fileurl = $my_address . "/" . $filepath;
    $data = json_encode(array("type" => $filetype, "url" => $fileurl)) . "\n";
    socket_write($socket, $data, strlen($data));
    
    while ($response = socket_read($socket, 2048)) {
        echo $response;
    }
    
    socket_close($socket);
    return true;
}
?>