curl -i -u "304d341f-4b5b-41e6-a124-40a9dc82432c:FA3gWEI4cjBE" \
-F training_data=@nlc/ipa-web-nlc_training.csv \
-F training_metadata="{\"language\":\"en\",\"name\":\"demo_ipa\"}" \
"https://gateway.watsonplatform.net/natural-language-classifier/api/v1/classifiers"
