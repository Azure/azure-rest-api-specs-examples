```javascript
const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for List all ResourceHealthMetadata for all sites in the subscription.
 *
 * @summary Description for List all ResourceHealthMetadata for all sites in the subscription.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ListResourceHealthMetadataBySubscription.json
 */
async function listResourceHealthMetadataForASubscription() {
  const subscriptionId = "4adb32ad-8327-4cbb-b775-b84b4465bb38";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.resourceHealthMetadataOperations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listResourceHealthMetadataForASubscription().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appservice_12.0.0/sdk/appservice/arm-appservice/README.md) on how to add the SDK to your project and authenticate.
