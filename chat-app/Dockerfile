# STAGE 1: Build
FROM --platform=linux/amd64 node:18-alpine as build
ARG config=production

WORKDIR /app-front
COPY package.json package-lock.json ./

RUN npm i --legacy-peer-deps

COPY . .

RUN echo "Building with config: ${config}"
RUN npm run build:${config}
RUN mkdir static

RUN cp -r dist/front/* static/

# STAGE 2: Run
FROM --platform=linux/amd64 nginx:1.13.12-alpine

# tzdata for timzone
RUN apk add --no-cache tzdata

# timezone env with default
ENV TZ=America/Argentina/Cordoba

COPY --from=build /app-front/static /usr/share/nginx/html
COPY config/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
