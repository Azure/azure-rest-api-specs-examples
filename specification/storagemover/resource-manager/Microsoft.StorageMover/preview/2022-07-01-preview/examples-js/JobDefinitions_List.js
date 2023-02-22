const { StorageMoverClient } = require("@azure/arm-storagemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all Job Definitions in a Project.
 *
 * @summary Lists all Job Definitions in a Project.
 * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2022-07-01-preview/examples/JobDefinitions_List.json
 */
async function jobDefinitionsList() {
  const subscriptionId =
    process.env["STORAGEMOVER_SUBSCRIPTION_ID"] || "11111111-2222-3333-4444-555555555555";
  const resourceGroupName = process.env["STORAGEMOVER_RESOURCE_GROUP"] || "examples-rg";
  const storageMoverName = "examples-storageMoverName";
  const projectName = "examples-projectName";
  const credential = new DefaultAzureCredential();
  const client = new StorageMoverClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.jobDefinitions.list(
    resourceGroupName,
    storageMoverName,
    projectName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
