Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appplatform_2.0.0/sdk/appplatform/arm-appplatform/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update custom domain of one lifecycle application.
 *
 * @summary Create or update custom domain of one lifecycle application.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/CustomDomains_CreateOrUpdate.json
 */
async function customDomainsCreateOrUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const appName = "myapp";
  const domainName = "mydomain.com";
  const domainResource = {
    properties: {
      certName: "mycert",
      thumbprint: "934367bf1c97033f877db0f15cb1b586957d3133",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.customDomains.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    appName,
    domainName,
    domainResource
  );
  console.log(result);
}

customDomainsCreateOrUpdate().catch(console.error);
```
