const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a CGProfile if it doesn't exist or update an existing CGProfile.
 *
 * @summary create a CGProfile if it doesn't exist or update an existing CGProfile.
 * x-ms-original-file: 2026-06-01-preview/ContainerGroupProfileCreateOrUpdate_CreateConfidential.json
 */
async function confidentialContainerGroupProfile() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const result = await client.cgProfile.createOrUpdate("demo", "demo1", {
    location: "westeurope",
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
        securityContext: { capabilities: { add: ["CAP_NET_ADMIN"] }, privileged: false },
      },
    ],
    imageRegistryCredentials: [],
    ipAddress: { type: "Public", ports: [{ port: 8000, protocol: "TCP" }] },
    osType: "Linux",
    sku: "Confidential",
    zones: ["1"],
  });
  console.log(result);
}
