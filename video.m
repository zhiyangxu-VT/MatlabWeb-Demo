function result = image(filename)

video = VideoReader('filename');
result = get(video)
response = jsonencode(result);
disp(response);