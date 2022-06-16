const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

async function getMaintenanceConfiguration() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const configName = "default";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.maintenanceConfigurations.get(
    resourceGroupName,
    resourceName,
    configName
  );
  console.log(result);
}

getMaintenanceConfiguration().catch(console.error);
