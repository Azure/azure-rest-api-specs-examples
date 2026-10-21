const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a Blob Access Point configuration.
 *
 * @summary creates or updates a Blob Access Point configuration.
 * x-ms-original-file: 2026-09-01/BlobAccessPointConfigurations_Create.json
 */
async function createBlobAccessPointConfiguration() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.blobAccessPointConfigurations.create(
    "testrg",
    "teststorageaccount",
    "testaccesspointconfig",
    {
      location: "eastus",
      properties: {
        state: "Active",
        description: "NetApp ONTAP S3-compatible source",
        source: {
          sourceType: "NetAppOntap",
          connection: {
            connectionType: "Endpoint",
            endpoint: "https://netapp.contoso.com/bucket1",
            tlsVerification: "Perform",
          },
          auth: {
            authType: "AccessKey",
            accessKeyId: "example-access-key",
            secretAccessKey: "example-secret-key",
            signingRegion: "us-east-1",
          },
        },
      },
    },
  );
  console.log(result);
}
