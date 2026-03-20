const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a managed cluster.
 *
 * @summary creates or updates a managed cluster.
 * x-ms-original-file: 2026-01-01/ManagedClustersCreate_AzureServiceMesh.json
 */
async function createOrUpdateManagedClusterWithAzureServiceMesh() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedClusters.createOrUpdate("rg1", "clustername1", {
    location: "location1",
    addonProfiles: {
      azureKeyvaultSecretsProvider: {
        config: { enableSecretRotation: "true", rotationPollInterval: "2m" },
        enabled: true,
      },
    },
    agentPoolProfiles: [
      {
        name: "nodepool1",
        type: "VirtualMachineScaleSets",
        count: 3,
        enableNodePublicIP: true,
        mode: "System",
        osType: "Linux",
        vmSize: "Standard_DS2_v2",
      },
    ],
    autoScalerProfile: { scaleDownDelayAfterAdd: "15m", scanInterval: "20s" },
    diskEncryptionSetID:
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des",
    dnsPrefix: "dnsprefix1",
    enableRbac: true,
    kubernetesVersion: "",
    linuxProfile: { adminUsername: "azureuser", ssh: { publicKeys: [{ keyData: "keydata" }] } },
    networkProfile: {
      loadBalancerProfile: { managedOutboundIPs: { count: 2 } },
      loadBalancerSku: "standard",
      outboundType: "loadBalancer",
    },
    serviceMeshProfile: {
      istio: {
        certificateAuthority: {
          plugin: {
            certChainObjectName: "cert-chain",
            certObjectName: "ca-cert",
            keyObjectName: "ca-key",
            keyVaultId:
              "/subscriptions/854c9ddb-fe9e-4aea-8d58-99ed88282881/resourceGroups/ddama-test/providers/Microsoft.KeyVault/vaults/my-akv",
            rootCertObjectName: "root-cert",
          },
        },
        components: {
          egressGateways: [
            {
              name: "test-istio-egress",
              enabled: true,
              gatewayConfigurationName: "test-gateway-configuration",
            },
          ],
          ingressGateways: [{ enabled: true, mode: "Internal" }],
        },
      },
      mode: "Istio",
    },
    servicePrincipalProfile: { clientId: "clientid", secret: "secret" },
    windowsProfile: { adminPassword: "replacePassword1234$", adminUsername: "azureuser" },
    sku: { name: "Basic", tier: "Free" },
    tags: { archv2: "", tier: "production" },
  });
  console.log(result);
}
