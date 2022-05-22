Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appplatform_2.0.0/sdk/appplatform/arm-appplatform/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a KPack build result log download URL.
 *
 * @summary Get a KPack build result log download URL.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/BuildService_GetBuildResultLog.json
 */
async function buildServiceGetBuildResultLog() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const buildServiceName = "default";
  const buildName = "mybuild";
  const buildResultName = "123";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.buildServiceOperations.getBuildResultLog(
    resourceGroupName,
    serviceName,
    buildServiceName,
    buildName,
    buildResultName
  );
  console.log(result);
}

buildServiceGetBuildResultLog().catch(console.error);
```
