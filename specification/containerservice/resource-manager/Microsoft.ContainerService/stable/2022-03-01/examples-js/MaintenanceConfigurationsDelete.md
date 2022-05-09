Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-containerservice_16.1.0-beta.1/sdk/containerservice/arm-containerservice/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a maintenance configuration.
 *
 * @summary Deletes a maintenance configuration.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-03-01/examples/MaintenanceConfigurationsDelete.json
 */
async function deleteMaintenanceConfiguration() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const configName = "default";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.maintenanceConfigurations.delete(
    resourceGroupName,
    resourceName,
    configName
  );
  console.log(result);
}

deleteMaintenanceConfiguration().catch(console.error);
```
