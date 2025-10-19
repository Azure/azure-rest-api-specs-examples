const { StorageMoverClient } = require("@azure/arm-storagemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a Job Definition resource, which contains configuration for a single unit of managed data transfer.
 *
 * @summary creates or updates a Job Definition resource, which contains configuration for a single unit of managed data transfer.
 * x-ms-original-file: 2025-07-01/JobDefinitions_CreateOrUpdate_CloudToCloud.json
 */
async function jobDefinitionsCreateOrUpdateCloudToCloud() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "60bcfc77-6589-4da2-b7fd-f9ec9322cf95";
  const client = new StorageMoverClient(credential, subscriptionId);
  const result = await client.jobDefinitions.createOrUpdate(
    "examples-rg",
    "examples-storageMoverName",
    "examples-projectName",
    "examples-jobDefinitionName",
    {
      properties: {
        description: "Example Job Definition Description",
        copyMode: "Additive",
        jobType: "CloudToCloud",
        sourceName: "examples-sourceEndpointName",
        sourceSubpath: "/",
        targetName: "examples-targetEndpointName",
        targetSubpath: "/",
        agentName: "dummy-agent",
      },
    },
  );
  console.log(result);
}
