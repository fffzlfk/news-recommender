FROM node:lts as dependencies
WORKDIR /frontend
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile

FROM node:lts as builder
WORKDIR /frontend
COPY . .
COPY --from=dependencies /frontend/node_modules ./node_modules
RUN yarn build

FROM node:lts as runner
WORKDIR /frontend
ENV NODE_ENV production
# If you are using a custom next.config.js file, uncomment this line.
# COPY --from=builder /frontend/next.config.js ./
COPY --from=builder /frontend/public ./public
COPY --from=builder /frontend/.next ./.next
COPY --from=builder /frontend/node_modules ./node_modules
COPY --from=builder /frontend/package.json ./package.json

EXPOSE 3000
CMD ["yarn", "start"]
