const { StorageMoverClient } = require("@azure/arm-storagemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an Endpoint resource, which represents a data transfer source or destination.
 *
 * @summary creates or updates an Endpoint resource, which represents a data transfer source or destination.
 * x-ms-original-file: 2025-07-01/Endpoints_CreateOrUpdate_SmbMount.json
 */
async function endpointsCreateOrUpdateSmbMount() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "60bcfc77-6589-4da2-b7fd-f9ec9322cf95";
  const client = new StorageMoverClient(credential, subscriptionId);
  const result = await client.endpoints.createOrUpdate(
    "examples-rg",
    "examples-storageMoverName",
    "examples-endpointName",
    {
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
    },
  );
  console.log(result);
}
