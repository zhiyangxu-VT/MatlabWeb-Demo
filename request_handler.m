t = tcpip('0.0.0.0', 3000, 'NetworkRole', 'server');
t.OutputBufferSize = 9999999999;

while(1)
    fopen(t);
    
    data = jsondecode(fscanf(t,'%s'));
    file_url = strcat('http://', data.url);
    disp(file_url)
    result = struct;
    
    content = webread(file_url);
    switch data.type
        case 'image'
            disp('image');
            result = imfinfo(file_url);
        case 'text'
            disp('text');
            result.type = 'text';
        case 'video'
            disp('video');
            result = get(content);
        case 'audio'
            disp('audio');
%             result = audioinfo(content);
    end

    disp(result);
    response = jsonencode(result);

    fwrite(t, response);
    fclose(t);
end
delete(t)
clear t;