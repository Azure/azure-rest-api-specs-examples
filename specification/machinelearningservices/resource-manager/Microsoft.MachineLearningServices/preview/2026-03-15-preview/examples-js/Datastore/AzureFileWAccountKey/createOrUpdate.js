const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update datastore.
 *
 * @summary create or update datastore.
 * x-ms-original-file: 2026-03-15-preview/Datastore/AzureFileWAccountKey/createOrUpdate.json
 */
async function createOrUpdateDatastoreAzureFileStoreWOrAccountKey() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.datastores.createOrUpdate(
    "test-rg",
    "my-aml-workspace",
    "string",
    {
      properties: {
        description: "string",
        accountName: "string",
        credentials: {
          credentialsType: "AccountKey",
          secrets: { key: "string", secretsType: "AccountKey" },
        },
        datastoreType: "AzureFile",
        endpoint: "string",
        fileShareName: "string",
        properties: {},
        tags: { string: "string" },
        protocol: "string",
      },
    },
    { skipValidation: false },
  );
  console.log(result);
}
