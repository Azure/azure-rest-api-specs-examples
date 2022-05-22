Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appplatform_2.0.0/sdk/appplatform/arm-appplatform/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to delete a Buildpack Binding
 *
 * @summary Operation to delete a Buildpack Binding
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/BuildpackBinding_Delete.json
 */
async function buildpackBindingDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const buildServiceName = "default";
  const builderName = "default";
  const buildpackBindingName = "myBuildpackBinding";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.buildpackBinding.beginDeleteAndWait(
    resourceGroupName,
    serviceName,
    buildServiceName,
    builderName,
    buildpackBindingName
  );
  console.log(result);
}

buildpackBindingDelete().catch(console.error);
```
