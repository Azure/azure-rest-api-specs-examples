const { StorageMoverClient } = require("@azure/arm-storagemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Job Definition resource.
 *
 * @summary Deletes a Job Definition resource.
 * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-10-01/examples/JobDefinitions_Delete.json
 */
async function projectsDelete() {
  const subscriptionId =
    process.env["STORAGEMOVER_SUBSCRIPTION_ID"] || "60bcfc77-6589-4da2-b7fd-f9ec9322cf95";
  const resourceGroupName = process.env["STORAGEMOVER_RESOURCE_GROUP"] || "examples-rg";
  const storageMoverName = "examples-storageMoverName";
  const projectName = "examples-projectName";
  const jobDefinitionName = "examples-jobDefinitionName";
  const credential = new DefaultAzureCredential();
  const client = new StorageMoverClient(credential, subscriptionId);
  const result = await client.jobDefinitions.beginDeleteAndWait(
    resourceGroupName,
    storageMoverName,
    projectName,
    jobDefinitionName
  );
  console.log(result);
}
