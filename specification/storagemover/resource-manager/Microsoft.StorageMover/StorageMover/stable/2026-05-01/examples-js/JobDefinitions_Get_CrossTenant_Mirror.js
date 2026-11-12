const { StorageMoverClient } = require("@azure/arm-storagemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets a Job Definition resource.
 *
 * @summary gets a Job Definition resource.
 * x-ms-original-file: 2026-05-01/JobDefinitions_Get_CrossTenant_Mirror.json
 */
async function jobDefinitionsGetCrossTenantMirror() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "0a2b3c4d-5e6f-7081-92a3-b4c5d6e7f809";
  const client = new StorageMoverClient(credential, subscriptionId);
  const result = await client.jobDefinitions.get(
    "partner-rg",
    "partner-storageMover",
    "partner-projectName",
    "examples-jobDefinitionName",
  );
  console.log(result);
}
