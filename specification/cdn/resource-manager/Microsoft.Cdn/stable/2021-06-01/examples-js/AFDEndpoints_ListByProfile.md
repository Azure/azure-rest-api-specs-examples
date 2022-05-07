Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-cdn_7.0.0/sdk/cdn/arm-cdn/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists existing AzureFrontDoor endpoints.
 *
 * @summary Lists existing AzureFrontDoor endpoints.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDEndpoints_ListByProfile.json
 */
async function afdEndpointsListByProfile() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.afdEndpoints.listByProfile(resourceGroupName, profileName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

afdEndpointsListByProfile().catch(console.error);
```
