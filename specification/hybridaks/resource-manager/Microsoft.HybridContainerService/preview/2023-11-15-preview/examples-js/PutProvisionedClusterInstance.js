const { HybridContainerServiceClient } = require("@azure/arm-hybridcontainerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates the Hybrid AKS provisioned cluster instance
 *
 * @summary Creates the Hybrid AKS provisioned cluster instance
 * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2023-11-15-preview/examples/PutProvisionedClusterInstance.json
 */
async function putProvisionedClusterInstance() {
  const connectedClusterResourceUri =
    "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster";
  const provisionedClusterInstance = {
    extendedLocation: {
      name: "/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourcegroups/test-arcappliance-resgrp/providers/microsoft.extendedlocation/customlocations/testcustomlocation",
      type: "CustomLocation",
    },
    properties: {
      agentPoolProfiles: [
        {
          name: "default-nodepool-1",
          count: 1,
          osType: "Linux",
          vmSize: "Standard_A4_v2",
        },
      ],
      cloudProviderProfile: {
        infraNetworkProfile: {
          vnetSubnetIds: [
            "/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.AzureStackHCI/logicalNetworks/test-vnet-static",
          ],
        },
      },
      controlPlane: {
        count: 1,
        linuxProfile: {
          ssh: {
            publicKeys: [{ keyData: "ssh-rsa AAAAB1NzaC1yc2EAAAADAQABAAACAQCY......" }],
          },
        },
        osType: "Linux",
        vmSize: "Standard_A4_v2",
      },
      kubernetesVersion: "v1.20.5",
      licenseProfile: { azureHybridBenefit: "NotApplicable" },
      linuxProfile: {
        ssh: {
          publicKeys: [{ keyData: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCY......." }],
        },
      },
      networkProfile: { networkPolicy: "calico", podCidr: "10.244.0.0/16" },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new HybridContainerServiceClient(credential);
  const result = await client.provisionedClusterInstances.beginCreateOrUpdateAndWait(
    connectedClusterResourceUri,
    provisionedClusterInstance
  );
  console.log(result);
}
