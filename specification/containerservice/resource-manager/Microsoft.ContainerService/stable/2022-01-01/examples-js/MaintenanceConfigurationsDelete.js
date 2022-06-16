const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

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
