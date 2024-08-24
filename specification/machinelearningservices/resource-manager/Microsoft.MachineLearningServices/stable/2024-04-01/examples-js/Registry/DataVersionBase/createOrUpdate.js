const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update version.
 *
 * @summary Create or update version.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Registry/DataVersionBase/createOrUpdate.json
 */
async function createOrUpdateRegistryDataVersionBase() {
  const subscriptionId =
    process.env["MACHINELEARNING_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["MACHINELEARNING_RESOURCE_GROUP"] || "test-rg";
  const registryName = "registryName";
  const name = "string";
  const version = "string";
  const body = {
    properties: {
      description: "string",
      dataType: "mltable",
      dataUri: "string",
      isAnonymous: false,
      isArchived: false,
      properties: { string: "string" },
      referencedUris: ["string"],
      tags: { string: "string" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.registryDataVersions.beginCreateOrUpdateAndWait(
    resourceGroupName,
    registryName,
    name,
    version,
    body,
  );
  console.log(result);
}
