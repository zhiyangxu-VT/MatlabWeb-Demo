<?php
$uploadOk = 1;
$imageFileType = pathinfo(basename($_FILES["imageFile"]["name"]),PATHINFO_EXTENSION);
// Check if image file is a actual image or fake image
if(isset($_POST["submit"])) {
    $check = getimagesize($_FILES["imageFile"]["tmp_name"]);
    if($check !== false) {
        echo "File is an image - " . $check["mime"] . ".";
        $uploadOk = 1;
    } else {
        echo "File is not an image.";
        $uploadOk = 0;
    }
}

// Check file size
if ($_FILES["imageFile"]["size"] > 500000) {
    echo "Sorry, your file is too large.";
    $uploadOk = 0;
}
// Allow certain file formats
if($imageFileType != "jpg" && $imageFileType != "png" && $imageFileType != "jpeg"
&& $imageFileType != "gif" ) {
    echo "Sorry, only JPG, JPEG, PNG & GIF files are allowed.";
    $uploadOk = 0;
}

// Check if $uploadOk is set to 0 by an error
if ($uploadOk == 0) {
    echo "Sorry, your file was not uploaded.";
// if everything is ok, try to upload file
} else {
    $date = new DateTime();
    $target_dir = "uploaded_img/";
    $target_file = $target_dir . "img-" . $date->getTimestamp() . "." . $imageFileType;
    
    if(!file_exists($target_dir)) {
        mkdir($target_dir);
    }
    while (file_exists($target_file)) {
        $target_file = $target_dir . "img-" . DateTime::getTimestamp();
    }
    if (move_uploaded_file($_FILES["imageFile"]["tmp_name"], $target_file)) {
        matlab_analysis($target_file);
    } else {
        echo "Sorry, there was an error uploading your file.";
    }
}


function matlab_analysis($filepath) {
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
    $data = $my_address . "/" . $filepath . "\n";
    socket_write($socket, $data, strlen($data));
    
    while ($response = socket_read($socket, 2048)) {
        echo $response;
    }
    
    socket_close($socket);
    return true;
}
?>