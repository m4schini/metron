Events
---

| Status              | Channel                 | Payload               | Description                                     |
|---------------------|-------------------------|-----------------------|-------------------------------------------------|
| `implemented`       | account.scan.requested  | `"username"`          | a scan of specified account should be performed |
| `implemented`       | account.scanned         | `{account}`           | the result of an account scan                   |
| `implemented`       | video.scan.requested    | `"url"`               | a scan of specified video should be performed   |
| `implemented`       | video.scanned           | `{video}`             | the result of an video scan                     |
| `specification`     | video.discovered        | `{username, id}`      | a video not known to the system                 |
| `specification`     | video.source.downloaded | `{username, id, url}` | a video source not known to the system          |
