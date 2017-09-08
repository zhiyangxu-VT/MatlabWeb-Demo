function result = audio_analyse(filename)

result = audioinfo(filename)
result = rmfield(result, 'Title')
result = rmfield(result, 'Comment')
result = rmfield(result, 'Artist')