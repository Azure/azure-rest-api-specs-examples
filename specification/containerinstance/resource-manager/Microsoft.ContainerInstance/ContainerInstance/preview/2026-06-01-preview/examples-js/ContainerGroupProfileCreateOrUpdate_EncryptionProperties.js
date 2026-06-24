const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a CGProfile if it doesn't exist or update an existing CGProfile.
 *
 * @summary create a CGProfile if it doesn't exist or update an existing CGProfile.
 * x-ms-original-file: 2026-06-01-preview/ContainerGroupProfileCreateOrUpdate_EncryptionProperties.json
 */
async function containerGroupProfileWithEncryptionProperties() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const result = await client.cgProfile.createOrUpdate("demo", "demo1", {
    location: "eastus2",
    containers: [
      {
        name: "demo1",
        command: [],
        environmentVariables: [],
        image: "nginx",
        ports: [{ port: 80 }],
        resources: { requests: { cpu: 1, memoryInGB: 1.5 } },
      },
    ],
    encryptionProperties: {
      identity:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/test-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/container-group-identity",
      keyName: "test-key",
      keyVersion: "<key version>",
      vaultBaseUrl: "https://testkeyvault.vault.azure.net",
    },
    imageRegistryCredentials: [],
    ipAddress: { type: "Public", ports: [{ port: 80, protocol: "TCP" }] },
    osType: "Linux",
    zones: ["1"],
  });
  console.log(result);
}
