t = tcpip('0.0.0.0', 3000, 'NetworkRole', 'server');
t.OutputBufferSize = 9999999999;

while(1)
    disp("Waiting for connection");
    drawnow
    fopen(t);
    disp("connected");
    drawnow

    % while(t.BytesAvailable<=0)
    %     drawnow
    % end
    % disp("Data recieved");
    data = fscanf(t,'%s');
    disp(data);
    img_file_url = char("http://" + data);
    result = imfinfo(img_file_url);
    disp(result);
    response = jsonencode(result);

    fwrite(t, response);
    fclose(t);
end
delete(t)
clear t;