const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update datastore.
 *
 * @summary Create or update datastore.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Datastore/AzureDataLakeGen1WServicePrincipal/createOrUpdate.json
 */
async function createOrUpdateDatastoreAzureDataLakeGen1WOrServicePrincipal() {
  const subscriptionId =
    process.env["MACHINELEARNING_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["MACHINELEARNING_RESOURCE_GROUP"] || "test-rg";
  const workspaceName = "my-aml-workspace";
  const name = "string";
  const skipValidation = false;
  const body = {
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
  };
  const options = { skipValidation };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.datastores.createOrUpdate(
    resourceGroupName,
    workspaceName,
    name,
    body,
    options,
  );
  console.log(result);
}
