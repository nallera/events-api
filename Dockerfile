FROM

# Copy custom commands
ADD /.ci/ /commands/
RUN chmod a+x /commands/*

