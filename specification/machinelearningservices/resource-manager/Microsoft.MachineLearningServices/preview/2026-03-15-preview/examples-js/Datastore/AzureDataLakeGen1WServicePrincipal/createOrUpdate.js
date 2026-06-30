const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update datastore.
 *
 * @summary create or update datastore.
 * x-ms-original-file: 2026-03-15-preview/Datastore/AzureDataLakeGen1WServicePrincipal/createOrUpdate.json
 */
async function createOrUpdateDatastoreAzureDataLakeGen1WOrServicePrincipal() {
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
        credentials: {
          authorityUrl: "string",
          clientId: "00000000-1111-2222-3333-444444444444",
          credentialsType: "ServicePrincipal",
          resourceUrl: "string",
          secrets: { clientSecret: "string", secretsType: "ServicePrincipal" },
          tenantId: "00000000-1111-2222-3333-444444444444",
        },
        datastoreType: "AzureDataLakeGen1",
        properties: {},
        storeName: "string",
        tags: { string: "string" },
      },
    },
    { skipValidation: false },
  );
  console.log(result);
}
