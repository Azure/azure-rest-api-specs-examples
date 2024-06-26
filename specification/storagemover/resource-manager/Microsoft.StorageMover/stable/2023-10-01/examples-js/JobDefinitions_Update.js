const { StorageMoverClient } = require("@azure/arm-storagemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates properties for a Job Definition resource. Properties not specified in the request body will be unchanged.
 *
 * @summary Updates properties for a Job Definition resource. Properties not specified in the request body will be unchanged.
 * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-10-01/examples/JobDefinitions_Update.json
 */
async function jobDefinitionsUpdate() {
  const subscriptionId =
    process.env["STORAGEMOVER_SUBSCRIPTION_ID"] || "60bcfc77-6589-4da2-b7fd-f9ec9322cf95";
  const resourceGroupName = process.env["STORAGEMOVER_RESOURCE_GROUP"] || "examples-rg";
  const storageMoverName = "examples-storageMoverName";
  const projectName = "examples-projectName";
  const jobDefinitionName = "examples-jobDefinitionName";
  const jobDefinition = {
    description: "Updated Job Definition Description",
    agentName: "updatedAgentName",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageMoverClient(credential, subscriptionId);
  const result = await client.jobDefinitions.update(
    resourceGroupName,
    storageMoverName,
    projectName,
    jobDefinitionName,
    jobDefinition
  );
  console.log(result);
}
