FROM nginx:mainline-alpine
COPY index.html /usr/share/nginx/html
COPY nginx.conf /etc/nginx/
EXPOSE 8081
CMD ["nginx", "-g", "daemon off;"]
