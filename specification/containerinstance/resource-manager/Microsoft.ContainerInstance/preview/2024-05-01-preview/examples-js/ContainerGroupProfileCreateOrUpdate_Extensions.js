const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update container group profiles with specified configurations.
 *
 * @summary Create or update container group profiles with specified configurations.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/preview/2024-05-01-preview/examples/ContainerGroupProfileCreateOrUpdate_Extensions.json
 */
async function containerGroupProfileCreateWithExtensions() {
  const subscriptionId =
    process.env["CONTAINERINSTANCE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERINSTANCE_RESOURCE_GROUP"] || "demo";
  const containerGroupProfileName = "demo1";
  const containerGroupProfile = {
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
    location: "eastus2",
    osType: "Linux",
    zones: ["1"],
    extensions: [
      {
        name: "kube-proxy",
        extensionType: "kube-proxy",
        protectedSettings: { kubeConfig: "<kubeconfig encoded string>" },
        settings: { clusterCidr: "10.240.0.0/16", kubeVersion: "v1.9.10" },
        version: "1.0",
      },
      {
        name: "vk-realtime-metrics",
        extensionType: "realtime-metrics",
        version: "1.0",
      },
    ],
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
