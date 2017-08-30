t = tcpip('0.0.0.0', 3000, 'NetworkRole', 'server', 'InputBufferSize', 1024);

while(1)
    fopen(t);
    connect = true;
    while(connect)
        while(t.BytesAvailable<=0)
            drawnow
        end
        data = fread(t, t.BytesAvailable);
        %disp(data);
        result = imfinfo("http://" + data);
        %disp(result["FileName"]);
        response = jsonencode(result);
        
        fwrite(t, response);
    end
    fclose(t);
end