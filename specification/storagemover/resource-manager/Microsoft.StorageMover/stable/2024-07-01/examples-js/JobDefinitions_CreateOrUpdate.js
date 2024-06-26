const { StorageMoverClient } = require("@azure/arm-storagemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Job Definition resource, which contains configuration for a single unit of managed data transfer.
 *
 * @summary Creates or updates a Job Definition resource, which contains configuration for a single unit of managed data transfer.
 * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2024-07-01/examples/JobDefinitions_CreateOrUpdate.json
 */
async function jobDefinitionsCreateOrUpdate() {
  const subscriptionId =
    process.env["STORAGEMOVER_SUBSCRIPTION_ID"] || "60bcfc77-6589-4da2-b7fd-f9ec9322cf95";
  const resourceGroupName = process.env["STORAGEMOVER_RESOURCE_GROUP"] || "examples-rg";
  const storageMoverName = "examples-storageMoverName";
  const projectName = "examples-projectName";
  const jobDefinitionName = "examples-jobDefinitionName";
  const jobDefinition = {
    description: "Example Job Definition Description",
    agentName: "migration-agent",
    copyMode: "Additive",
    sourceName: "examples-sourceEndpointName",
    sourceSubpath: "/",
    targetName: "examples-targetEndpointName",
    targetSubpath: "/",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageMoverClient(credential, subscriptionId);
  const result = await client.jobDefinitions.createOrUpdate(
    resourceGroupName,
    storageMoverName,
    projectName,
    jobDefinitionName,
    jobDefinition,
  );
  console.log(result);
}
