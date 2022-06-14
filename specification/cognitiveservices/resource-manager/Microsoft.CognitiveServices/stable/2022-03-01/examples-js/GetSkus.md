```javascript
const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of Microsoft.CognitiveServices SKUs available for your Subscription.
 *
 * @summary Gets the list of Microsoft.CognitiveServices SKUs available for your Subscription.
 * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/GetSkus.json
 */
async function regenerateKeys() {
  const subscriptionId = "f1c637e4-72ec-4f89-8d2b-0f933c036002";
  const credential = new DefaultAzureCredential();
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.resourceSkus.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

regenerateKeys().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-cognitiveservices_7.2.0/sdk/cognitiveservices/arm-cognitiveservices/README.md) on how to add the SDK to your project and authenticate.
