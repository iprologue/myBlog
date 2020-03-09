FROM scratch

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $HOME/Developer/myBlog
COPY . $HOME/Developer/myBlog

EXPOSE 8000
CMD ["./myBlog"]