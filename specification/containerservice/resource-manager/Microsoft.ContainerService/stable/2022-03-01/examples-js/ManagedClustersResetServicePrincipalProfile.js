const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This action cannot be performed on a cluster that is not using a service principal
 *
 * @summary This action cannot be performed on a cluster that is not using a service principal
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-03-01/examples/ManagedClustersResetServicePrincipalProfile.json
 */
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
