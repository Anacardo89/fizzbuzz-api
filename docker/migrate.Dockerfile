FROM migrate/migrate:latest
WORKDIR /migrations
COPY migrations ./migrations
COPY scripts/ ./scripts/

ENTRYPOINT ["/scripts/wait-for-db.sh"]