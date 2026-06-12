const { StorageMoverClient } = require("@azure/arm-storagemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an Endpoint resource, which represents a data transfer source or destination.
 *
 * @summary creates or updates an Endpoint resource, which represents a data transfer source or destination.
 * x-ms-original-file: 2025-12-01/Endpoints_CreateOrUpdate_S3WithHMAC.json
 */
async function endpointsCreateOrUpdateS3WithHmac() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "60bcfc77-6589-4da2-b7fd-f9ec9322cf95";
  const client = new StorageMoverClient(credential, subscriptionId);
  const result = await client.endpoints.createOrUpdate(
    "examples-rg",
    "examples-storageMoverName",
    "examples-endpointName",
    {
      properties: {
        description: "Example S3WithHmac Endpoint Description",
        credentials: {
          accessKeyUri: "https://examples-azureKeyVault.vault.azure.net/secrets/examples-access",
          secretKeyUri: "https://examples-azureKeyVault.vault.azure.net/secrets/examples-secret",
          type: "AzureKeyVaultS3WithHMAC",
        },
        sourceUri: "https://examples-bucket.s3.amazonaws.com/prefix/",
        sourceType: "GCS",
        endpointType: "S3WithHMAC",
        endpointKind: "Source",
      },
    },
  );
  console.log(result);
}
