Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appplatform_2.0.0/sdk/appplatform/arm-appplatform/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to update an exiting App.
 *
 * @summary Operation to update an exiting App.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Apps_Update.json
 */
async function appsUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const appName = "myapp";
  const appResource = {
    identity: {
      type: "SystemAssigned",
      principalId: undefined,
      tenantId: undefined,
    },
    location: "eastus",
    properties: {
      enableEndToEndTLS: false,
      fqdn: "myapp.mydomain.com",
      httpsOnly: false,
      persistentDisk: { mountPath: "/mypersistentdisk", sizeInGB: 2 },
      public: true,
      temporaryDisk: { mountPath: "/mytemporarydisk", sizeInGB: 2 },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.apps.beginUpdateAndWait(
    resourceGroupName,
    serviceName,
    appName,
    appResource
  );
  console.log(result);
}

appsUpdate().catch(console.error);
```
