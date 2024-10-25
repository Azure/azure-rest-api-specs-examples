const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update container group profiles with specified configurations.
 *
 * @summary Create or update container group profiles with specified configurations.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/preview/2024-05-01-preview/examples/ContainerGroupProfileCreateOrUpdate_CreateConfidential.json
 */
async function confidentialContainerGroupProfile() {
  const subscriptionId =
    process.env["CONTAINERINSTANCE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERINSTANCE_RESOURCE_GROUP"] || "demo";
  const containerGroupProfileName = "demo1";
  const containerGroupProfile = {
    confidentialComputeProperties: {
      ccePolicy:
        "eyJhbGxvd19hbGwiOiB0cnVlLCAiY29udGFpbmVycyI6IHsibGVuZ3RoIjogMCwgImVsZW1lbnRzIjogbnVsbH19",
    },
    containers: [
      {
        name: "accdemo",
        command: [],
        environmentVariables: [],
        image: "confiimage",
        ports: [{ port: 8000 }],
        resources: { requests: { cpu: 1, memoryInGB: 1.5 } },
        securityContext: {
          capabilities: { add: ["CAP_NET_ADMIN"] },
          privileged: false,
        },
      },
    ],
    imageRegistryCredentials: [],
    ipAddress: { type: "Public", ports: [{ port: 8000, protocol: "TCP" }] },
    location: "westeurope",
    osType: "Linux",
    sku: "Confidential",
    zones: ["1"],
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const result = await client.containerGroupProfiles.createOrUpdate(
    resourceGroupName,
    containerGroupProfileName,
    containerGroupProfile,
  );
  console.log(result);
}
