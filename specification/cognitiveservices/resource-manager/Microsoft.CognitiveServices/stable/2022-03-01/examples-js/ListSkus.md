Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-cognitiveservices_7.1.0/sdk/cognitiveservices/arm-cognitiveservices/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List available SKUs for the requested Cognitive Services account
 *
 * @summary List available SKUs for the requested Cognitive Services account
 * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/ListSkus.json
 */
async function listSkUs() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const resourceGroupName = "myResourceGroup";
  const accountName = "myAccount";
  const credential = new DefaultAzureCredential();
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.accounts.listSkus(resourceGroupName, accountName);
  console.log(result);
}

listSkUs().catch(console.error);
```
