const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update container group profiles with specified configurations.
 *
 * @summary Create or update container group profiles with specified configurations.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/preview/2024-05-01-preview/examples/ContainerGroupProfilesCreateOrUpdate.json
 */
async function containerGroupProfilesCreateOrUpdate() {
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
        resources: {
          requests: { cpu: 1, gpu: { count: 1, sku: "K80" }, memoryInGB: 1.5 },
        },
        volumeMounts: [
          { name: "volume1", mountPath: "/mnt/volume1", readOnly: false },
          { name: "volume2", mountPath: "/mnt/volume2", readOnly: false },
          { name: "volume3", mountPath: "/mnt/volume3", readOnly: true },
        ],
      },
    ],
    diagnostics: {
      logAnalytics: {
        logType: "ContainerInsights",
        metadata: { podUuid: "test-metadata-value" },
        workspaceId: "workspaceid",
        workspaceKey: "workspaceKey",
        workspaceResourceId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg/providers/microsoft.operationalinsights/workspaces/workspace",
      },
    },
    imageRegistryCredentials: [],
    ipAddress: { type: "Public", ports: [{ port: 80, protocol: "TCP" }] },
    location: "west us",
    osType: "Linux",
    volumes: [
      {
        name: "volume1",
        azureFile: {
          shareName: "shareName",
          storageAccountKey: "accountKey",
          storageAccountName: "accountName",
        },
      },
      { name: "volume2", emptyDir: {} },
      {
        name: "volume3",
        secret: {
          secretKey1: "SecretValue1InBase64",
          secretKey2: "SecretValue2InBase64",
        },
      },
    ],
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
