## Slack Notice AWS Support

Notify slack of AWS support case with subject, status, submittedby, timecreated and url. It requires setting up [aws-actions/configure-aws-credentials](https://github.com/marketplace/actions/configure-aws-credentials-action-for-github-actions) and activating [Imcomming Webhooks](https://api.slack.com/messaging/webhooks).

![slack_sample](https://user-images.githubusercontent.com/65695538/155986836-8be88c7f-8b2f-4ee5-a3ee-d7572e37785c.png)

## Usage
Configure aws credentials before using this action.

```yaml
on:
  schedule:
    - cron: '0 1 * * 1-5'

permissions:
  id-token: write
  contents: write

jobs:
  notify-slack-aws-support:
    runs-on: ubuntu-latest
    steps:
      - name: Git clone the repository
        uses: actions/checkout@v1

      - name: Configure aws credentials
        uses: aws-actions/configure-aws-credentials@master
        with:
          role-to-assume: arn:aws:iam::123456789100:role/my-github-actions-role
          aws-region: ap-northeast-1
      
      - name:  Notify slack of aws support
        uses: sasamuku/slack_notice_aws_support@main
        with:
          include_resolved_cases: 'false'
          language: 'ja'
          webhook_url: 'http://hooks.slack.com/services/...'
```

See [action.yml](action.yml) for the full documentation for this action's inputs and outputs.

## AWS Credentials

The credentials used in GitHub Actions workflows must have permissions to [DescribeCases](https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeCases.html).

```yaml
    - Effect: 'Allow'
      Action:
        - 'support:DescribeCases'
      Resource: '*'
```

See [Access permissions for AWS Support](https://docs.aws.amazon.com/awssupport/latest/user/accessing-support.html) for more details.

## Imcomming webhooks
Activate [Incoming Webhooks](https://api.slack.com/messaging/webhooks) ant get the url similar to the following:
```
http://hooks.slack.com/services/...
```

## License Summary
This code is made available under the MIT license.
