const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the results of a command which has been run on the Managed Cluster.
 *
 * @summary Gets the results of a command which has been run on the Managed Cluster.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-08-02-preview/examples/RunCommandResultSucceed.json
 */
async function commandSucceedResult() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const commandId = "def7b3ea71bd4f7e9d226ddbc0f00ad9";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedClusters.getCommandResult(
    resourceGroupName,
    resourceName,
    commandId
  );
  console.log(result);
}

commandSucceedResult().catch(console.error);
