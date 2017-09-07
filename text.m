function result = image(filename)

fileID = fopen('nums1.txt','r');
result.some_content = fsacnf(fileID, '%s');
fclose(fileID);
response = jsonencode(result);
disp(response);