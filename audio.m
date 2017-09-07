function result = image(filename)

result = audioinfo(filename);
response = jsonencode(result);
disp(response);