const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This API allows you to upload content to a file
 *
 * @summary This API allows you to upload content to a file
 * x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2022-09-01-preview/examples/UploadFileForSubscription.json
 */
async function uploadFileForSubscription() {
  const subscriptionId = process.env["SUPPORT_SUBSCRIPTION_ID"] || "subid";
  const fileWorkspaceName = "testworkspaceName";
  const fileName = "test.txt";
  const uploadFile = {
    chunkIndex: 0,
    content:
      "iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAMAAAAoLQ9TAAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAAAgY0hSTQAAeiYAAICEAAD6AAAAgOgAAHUwAADqYAAAOpgAABd",
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const result = await client.files.upload(fileWorkspaceName, fileName, uploadFile);
  console.log(result);
}
