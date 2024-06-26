const { StorageMoverClient } = require("@azure/arm-storagemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Requests an Agent to start a new instance of this Job Definition, generating a new Job Run resource.
 *
 * @summary Requests an Agent to start a new instance of this Job Definition, generating a new Job Run resource.
 * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2023-07-01-preview/examples/JobDefinitions_StartJob.json
 */
async function jobDefinitionsStartJob() {
  const subscriptionId =
    process.env["STORAGEMOVER_SUBSCRIPTION_ID"] || "60bcfc77-6589-4da2-b7fd-f9ec9322cf95";
  const resourceGroupName = process.env["STORAGEMOVER_RESOURCE_GROUP"] || "examples-rg";
  const storageMoverName = "examples-storageMoverName";
  const projectName = "examples-projectName";
  const jobDefinitionName = "examples-jobDefinitionName";
  const credential = new DefaultAzureCredential();
  const client = new StorageMoverClient(credential, subscriptionId);
  const result = await client.jobDefinitions.startJob(
    resourceGroupName,
    storageMoverName,
    projectName,
    jobDefinitionName
  );
  console.log(result);
}
