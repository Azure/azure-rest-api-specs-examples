Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-cdn_7.0.1/sdk/cdn/arm-cdn/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new origin within the specified endpoint.
 *
 * @summary Creates a new origin within the specified endpoint.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Origins_Create.json
 */
async function originsCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const endpointName = "endpoint1";
  const originName = "www-someDomain-net";
  const origin = {
    enabled: true,
    hostName: "www.someDomain.net",
    httpPort: 80,
    httpsPort: 443,
    originHostHeader: "www.someDomain.net",
    priority: 1,
    privateLinkApprovalMessage: "Please approve the connection request for this Private Link",
    privateLinkLocation: "eastus",
    privateLinkResourceId:
      "/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.Network/privateLinkServices/pls1",
    weight: 50,
  };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.origins.beginCreateAndWait(
    resourceGroupName,
    profileName,
    endpointName,
    originName,
    origin
  );
  console.log(result);
}

originsCreate().catch(console.error);
```
