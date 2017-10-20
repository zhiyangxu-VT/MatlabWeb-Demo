function result = video_analyse(filename)

video = VideoReader(filename);
result = get(video);
result = rmfield(result, 'UserData');