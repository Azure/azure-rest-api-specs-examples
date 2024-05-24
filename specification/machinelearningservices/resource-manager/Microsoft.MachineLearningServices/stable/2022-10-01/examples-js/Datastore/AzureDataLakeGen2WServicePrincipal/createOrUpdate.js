const { AzureMachineLearningWorkspaces } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update datastore.
 *
 * @summary Create or update datastore.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Datastore/AzureDataLakeGen2WServicePrincipal/createOrUpdate.json
 */
async function createOrUpdateDatastoreAzureDataLakeGen2WOrServicePrincipal() {
  const subscriptionId =
    process.env["MACHINELEARNING_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["MACHINELEARNING_RESOURCE_GROUP"] || "test-rg";
  const workspaceName = "my-aml-workspace";
  const name = "string";
  const skipValidation = false;
  const body = {
    properties: {
      description: "string",
      accountName: "string",
      credentials: {
        authorityUrl: "string",
        clientId: "00000000-1111-2222-3333-444444444444",
        credentialsType: "ServicePrincipal",
        resourceUrl: "string",
        secrets: { clientSecret: "string", secretsType: "ServicePrincipal" },
        tenantId: "00000000-1111-2222-3333-444444444444",
      },
      datastoreType: "AzureDataLakeGen2",
      endpoint: "string",
      filesystem: "string",
      properties: {},
      tags: { string: "string" },
      protocol: "string",
    },
  };
  const options = { skipValidation };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningWorkspaces(credential, subscriptionId);
  const result = await client.datastores.createOrUpdate(
    resourceGroupName,
    workspaceName,
    name,
    body,
    options,
  );
  console.log(result);
}
