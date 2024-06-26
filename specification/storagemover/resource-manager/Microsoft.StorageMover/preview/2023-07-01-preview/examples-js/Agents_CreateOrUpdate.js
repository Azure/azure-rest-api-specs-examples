const { StorageMoverClient } = require("@azure/arm-storagemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an Agent resource, which references a hybrid compute machine that can run jobs.
 *
 * @summary Creates or updates an Agent resource, which references a hybrid compute machine that can run jobs.
 * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2023-07-01-preview/examples/Agents_CreateOrUpdate.json
 */
async function agentsCreateOrUpdate() {
  const subscriptionId =
    process.env["STORAGEMOVER_SUBSCRIPTION_ID"] || "60bcfc77-6589-4da2-b7fd-f9ec9322cf95";
  const resourceGroupName = process.env["STORAGEMOVER_RESOURCE_GROUP"] || "examples-rg";
  const storageMoverName = "examples-storageMoverName";
  const agentName = "examples-agentName";
  const agent = {
    description: "Example Agent Description",
    arcResourceId:
      "/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.HybridCompute/machines/examples-hybridComputeName",
    arcVmUuid: "3bb2c024-eba9-4d18-9e7a-1d772fcc5fe9",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageMoverClient(credential, subscriptionId);
  const result = await client.agents.createOrUpdate(
    resourceGroupName,
    storageMoverName,
    agentName,
    agent
  );
  console.log(result);
}
