Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appplatform_2.0.0/sdk/appplatform/arm-appplatform/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the config server.
 *
 * @summary Update the config server.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/ConfigServers_UpdatePut.json
 */
async function configServersUpdatePut() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const configServerResource = {
    properties: {
      configServer: {
        gitProperty: {
          label: "master",
          searchPaths: ["/"],
          uri: "https://github.com/fake-user/fake-repository.git",
        },
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.configServers.beginUpdatePutAndWait(
    resourceGroupName,
    serviceName,
    configServerResource
  );
  console.log(result);
}

configServersUpdatePut().catch(console.error);
```
