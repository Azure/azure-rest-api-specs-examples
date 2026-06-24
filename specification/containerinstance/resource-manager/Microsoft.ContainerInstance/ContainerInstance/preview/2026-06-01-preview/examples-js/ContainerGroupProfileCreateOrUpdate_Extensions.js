const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a CGProfile if it doesn't exist or update an existing CGProfile.
 *
 * @summary create a CGProfile if it doesn't exist or update an existing CGProfile.
 * x-ms-original-file: 2026-06-01-preview/ContainerGroupProfileCreateOrUpdate_Extensions.json
 */
async function containerGroupProfileCreateWithExtensions() {
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
    imageRegistryCredentials: [],
    ipAddress: { type: "Private", ports: [{ port: 80, protocol: "TCP" }] },
    osType: "Linux",
    extensions: [
      {
        name: "kube-proxy",
        extensionType: "kube-proxy",
        protectedSettings: { kubeConfig: "<kubeconfig encoded string>" },
        settings: { clusterCidr: "10.240.0.0/16", kubeVersion: "v1.9.10" },
        version: "1.0",
      },
      { name: "vk-realtime-metrics", extensionType: "realtime-metrics", version: "1.0" },
    ],
    zones: ["1"],
  });
  console.log(result);
}
