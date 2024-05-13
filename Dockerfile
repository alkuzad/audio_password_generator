FROM ubuntu:24.04

RUN apt update && apt install -y libmp3lame-dev

COPY --chmod=770 ./audio_password_generator /audio_password_generator
COPY ./sound /sound

ENTRYPOINT ["/audio_password_generator", "-dir", "/tmp/output", "-mp3", "-size"]

CMD ["-h"]