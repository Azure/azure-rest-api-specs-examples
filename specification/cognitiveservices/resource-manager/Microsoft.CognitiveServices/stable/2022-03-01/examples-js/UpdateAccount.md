```javascript
const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a Cognitive Services account
 *
 * @summary Updates a Cognitive Services account
 * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/UpdateAccount.json
 */
async function updateAccount() {
  const subscriptionId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const resourceGroupName = "bvttest";
  const accountName = "bingSearch";
  const account = { location: "global", sku: { name: "S2" } };
  const credential = new DefaultAzureCredential();
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.accounts.beginUpdateAndWait(resourceGroupName, accountName, account);
  console.log(result);
}

updateAccount().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-cognitiveservices_7.2.0/sdk/cognitiveservices/arm-cognitiveservices/README.md) on how to add the SDK to your project and authenticate.
