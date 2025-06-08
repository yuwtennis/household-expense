# household-expense

## Run

```shell
export GDRIVE_FOLDER_ID=$(gcloud secrets versions access --secret gdrive_folder_id latest)
export BUCKET_NAME=$(gcloud secrets versions access --secret bucket_name latest)
docker run \
-e GDRIVE_FOLDER_ID \
-e BUCKET_NAME \
-v ~/.config/:/root/.config/ \
-t household-expense:$(git rev-parse HEAD) \
app
```