const { StorageMoverClient } = require("@azure/arm-storagemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an Endpoint resource, which represents a data transfer source or destination.
 *
 * @summary Creates or updates an Endpoint resource, which represents a data transfer source or destination.
 * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2023-07-01-preview/examples/Endpoints_CreateOrUpdate_SmbMount.json
 */
async function endpointsCreateOrUpdateSmbMount() {
  const subscriptionId =
    process.env["STORAGEMOVER_SUBSCRIPTION_ID"] || "60bcfc77-6589-4da2-b7fd-f9ec9322cf95";
  const resourceGroupName = process.env["STORAGEMOVER_RESOURCE_GROUP"] || "examples-rg";
  const storageMoverName = "examples-storageMoverName";
  const endpointName = "examples-endpointName";
  const endpoint = {
    properties: {
      description: "Example SMB Mount Endpoint Description",
      credentials: {
        type: "AzureKeyVaultSmb",
        passwordUri: "https://examples-azureKeyVault.vault.azure.net/secrets/examples-password",
        usernameUri: "https://examples-azureKeyVault.vault.azure.net/secrets/examples-username",
      },
      endpointType: "SmbMount",
      host: "0.0.0.0",
      shareName: "examples-shareName",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageMoverClient(credential, subscriptionId);
  const result = await client.endpoints.createOrUpdate(
    resourceGroupName,
    storageMoverName,
    endpointName,
    endpoint
  );
  console.log(result);
}
