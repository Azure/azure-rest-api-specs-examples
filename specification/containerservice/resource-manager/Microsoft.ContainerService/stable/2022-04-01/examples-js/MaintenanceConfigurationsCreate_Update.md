```javascript
const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a maintenance configuration in the specified managed cluster.
 *
 * @summary Creates or updates a maintenance configuration in the specified managed cluster.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-04-01/examples/MaintenanceConfigurationsCreate_Update.json
 */
async function createOrUpdateMaintenanceConfiguration() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const configName = "default";
  const parameters = {
    notAllowedTime: [
      {
        end: new Date("2020-11-30T12:00:00Z"),
        start: new Date("2020-11-26T03:00:00Z"),
      },
    ],
    timeInWeek: [{ day: "Monday", hourSlots: [1, 2] }],
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.maintenanceConfigurations.createOrUpdate(
    resourceGroupName,
    resourceName,
    configName,
    parameters
  );
  console.log(result);
}

createOrUpdateMaintenanceConfiguration().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-containerservice_16.1.0/sdk/containerservice/arm-containerservice/README.md) on how to add the SDK to your project and authenticate.
