const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

async function resetServicePrincipalProfile() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const parameters = {
    clientId: "clientid",
    secret: "secret",
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedClusters.beginResetServicePrincipalProfileAndWait(
    resourceGroupName,
    resourceName,
    parameters
  );
  console.log(result);
}

resetServicePrincipalProfile().catch(console.error);
