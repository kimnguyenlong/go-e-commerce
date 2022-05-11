FROM nginx:latest   
RUN rm /etc/nginx/conf.d/default.conf
COPY main.conf /etc/nginx/conf.d
RUN mkdir /etc/nginx/ssl
RUN chmod 700 /etc/nginx/ssl
COPY ./ssl/cert.crt /etc/nginx/ssl
COPY ./ssl/pri.key /etc/nginx/ssl

EXPOSE 80 443
CMD ["nginx", "-g", "daemon off;"]
