const { AzureMachineLearningServicesManagementClient } = require("@azure/arm-machinelearning");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get blob reference SAS Uri.
 *
 * @summary Get blob reference SAS Uri.
 * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/DataReference/getBlobReferenceSAS.json
 */
async function getBlobReferenceSasDataReference() {
  const subscriptionId =
    process.env["MACHINELEARNING_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["MACHINELEARNING_RESOURCE_GROUP"] || "test-rg";
  const registryName = "registryName";
  const name = "string";
  const version = "string";
  const body = {
    assetId: "string",
    blobUri: "https://www.contoso.com/example",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMachineLearningServicesManagementClient(credential, subscriptionId);
  const result = await client.registryDataReferences.getBlobReferenceSAS(
    resourceGroupName,
    registryName,
    name,
    version,
    body,
  );
  console.log(result);
}
