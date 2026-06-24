const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a CGProfile if it doesn't exist or update an existing CGProfile.
 *
 * @summary create a CGProfile if it doesn't exist or update an existing CGProfile.
 * x-ms-original-file: 2026-06-01-preview/ContainerGroupProfilesCreateOrUpdate.json
 */
async function containerGroupProfilesCreateOrUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const result = await client.cgProfile.createOrUpdate("demo", "demo1", {
    location: "west us",
    containers: [
      {
        name: "demo1",
        command: [],
        environmentVariables: [],
        image: "nginx",
        ports: [{ port: 80 }],
        resources: { requests: { cpu: 1, gpu: { count: 1, sku: "K80" }, memoryInGB: 1.5 } },
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
        metadata: { "pod-uuid": "test-metadata-value" },
        workspaceId: "workspaceid",
        workspaceKey: "workspaceKey",
        workspaceResourceId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg/providers/microsoft.operationalinsights/workspaces/workspace",
      },
    },
    imageRegistryCredentials: [],
    ipAddress: { type: "Public", ports: [{ port: 80, protocol: "TCP" }] },
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
        secret: { secretKey1: "SecretValue1InBase64", secretKey2: "SecretValue2InBase64" },
      },
    ],
    zones: ["1"],
  });
  console.log(result);
}
