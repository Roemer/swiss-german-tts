FROM python:3.9-bullseye

SHELL ["/bin/bash", "-c"]

RUN apt update && apt install espeak-ng git -y

WORKDIR /prj

RUN git clone https://github.com/rhasspy/piper.git

WORKDIR /prj/piper/src/python

RUN python3 -m venv .venv \
    && source .venv/bin/activate \
    && pip3 install --upgrade pip \
    && pip3 install --upgrade wheel setuptools \
    && pip3 install -r requirements.txt

RUN ./build_monotonic_align.sh
