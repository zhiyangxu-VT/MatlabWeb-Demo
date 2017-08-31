function result = basic_info(filename)

result = imfinfo(filename);
response = jsonencode(result);
disp(response);