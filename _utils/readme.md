https://github.com/rhasspy/piper

docker build -t tts .
docker run --rm -it -v C:\git\data1.1:/data --shm-size 1G tts /bin/bash
source .venv/bin/activate

# PreProcess
python3 -m piper_train.preprocess --language de-ch --input-dir /data --output-dir /data/_out/ --dataset-format ljspeech --sample-rate 22050

# Finally, you can train:
# python3 -m piper_train --help
python3 -m piper_train \
    --dataset-dir /data/_out/ \
    --accelerator 'cpu' \
    --devices 1 \
    --batch-size 32 \
    --validation-split 0.05 \
    --num-test-examples 5 \
    --max_epochs 10000 \
    --precision 32

docker exec -it 7d7be35fd0f3 /bin/bash
cd /prj/piper/src/python
tensorboard --logdir /data/_out/lightning_logs