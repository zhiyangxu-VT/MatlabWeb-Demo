function result = image_analyse(filename)

result = imfinfo(filename);
result = rmfield(result, 'ColorTable');
disp(result);