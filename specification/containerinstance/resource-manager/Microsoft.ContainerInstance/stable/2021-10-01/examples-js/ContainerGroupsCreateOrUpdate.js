const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update container groups with specified configurations.
 *
 * @summary Create or update container groups with specified configurations.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2021-10-01/examples/ContainerGroupsCreateOrUpdate.json
 */
async function containerGroupsCreateOrUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "demo";
  const containerGroupName = "demo1";
  const containerGroup = {
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
        metadata: { testKey: "test-metadata-value" },
        workspaceId: "workspaceid",
        workspaceKey: "workspaceKey",
        workspaceResourceId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg/providers/microsoft.operationalinsights/workspaces/workspace",
      },
    },
    dnsConfig: {
      nameServers: ["1.1.1.1"],
      options: "ndots:2",
      searchDomains: "cluster.local svc.cluster.local",
    },
    identity: {
      type: "SystemAssigned, UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/00000000000000000000000000000000/resourceGroups/myResourceGroup/providers/MicrosoftManagedIdentity/userAssignedIdentities/identityName":
          {},
      },
    },
    imageRegistryCredentials: [],
    ipAddress: {
      type: "Public",
      dnsNameLabel: "dnsnamelabel1",
      dnsNameLabelReusePolicy: "Unsecure",
      ports: [{ port: 80, protocol: "TCP" }],
    },
    location: "west us",
    osType: "Linux",
    subnetIds: [
      {
        id: "[resourceId('Microsoft.Network/virtualNetworks/subnets', parameters('vnetName'), parameters('subnetName'))]",
      },
    ],
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
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const result = await client.containerGroups.beginCreateOrUpdateAndWait(
    resourceGroupName,
    containerGroupName,
    containerGroup
  );
  console.log(result);
}

containerGroupsCreateOrUpdate().catch(console.error);
