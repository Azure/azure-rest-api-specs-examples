const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

async function resetAadProfile() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const parameters = {
    clientAppID: "clientappid",
    serverAppID: "serverappid",
    serverAppSecret: "serverappsecret",
    tenantID: "tenantid",
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedClusters.beginResetAADProfileAndWait(
    resourceGroupName,
    resourceName,
    parameters
  );
  console.log(result);
}

resetAadProfile().catch(console.error);
