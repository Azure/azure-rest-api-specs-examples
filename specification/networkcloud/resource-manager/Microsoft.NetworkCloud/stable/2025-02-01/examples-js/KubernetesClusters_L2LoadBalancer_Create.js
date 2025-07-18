const { NetworkCloud } = require("@azure/arm-networkcloud");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create a new Kubernetes cluster or update the properties of the existing one.
 *
 * @summary Create a new Kubernetes cluster or update the properties of the existing one.
 * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/KubernetesClusters_L2LoadBalancer_Create.json
 */
async function createOrUpdateKubernetesClusterWithALayer2LoadBalancer() {
  const subscriptionId =
    process.env["NETWORKCLOUD_SUBSCRIPTION_ID"] || "123e4567-e89b-12d3-a456-426655440000";
  const resourceGroupName = process.env["NETWORKCLOUD_RESOURCE_GROUP"] || "resourceGroupName";
  const kubernetesClusterName = "kubernetesClusterName";
  const kubernetesClusterParameters = {
    aadConfiguration: {
      adminGroupObjectIds: ["ffffffff-ffff-ffff-ffff-ffffffffffff"],
    },
    administratorConfiguration: {
      adminUsername: "azure",
      sshPublicKeys: [
        {
          keyData:
            "ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm",
        },
      ],
    },
    controlPlaneNodeConfiguration: {
      administratorConfiguration: {
        adminUsername: "azure",
        sshPublicKeys: [
          {
            keyData:
              "ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm",
          },
        ],
      },
      availabilityZones: ["1", "2", "3"],
      count: 3,
      vmSkuName: "NC_G6_28_v1",
    },
    extendedLocation: {
      name: "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName",
      type: "CustomLocation",
    },
    initialAgentPoolConfigurations: [
      {
        name: "SystemPool-1",
        administratorConfiguration: {
          adminUsername: "azure",
          sshPublicKeys: [
            {
              keyData:
                "ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm",
            },
          ],
        },
        agentOptions: { hugepagesCount: 96, hugepagesSize: "1G" },
        attachedNetworkConfiguration: {
          l2Networks: [
            {
              networkId:
                "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l2Networks/l2NetworkName",
              pluginType: "DPDK",
            },
          ],
          l3Networks: [
            {
              ipamEnabled: "False",
              networkId:
                "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName",
              pluginType: "SRIOV",
            },
          ],
          trunkedNetworks: [
            {
              networkId:
                "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/trunkedNetworks/trunkedNetworkName",
              pluginType: "MACVLAN",
            },
          ],
        },
        availabilityZones: ["1", "2", "3"],
        count: 3,
        labels: [{ key: "kubernetes.label", value: "true" }],
        mode: "System",
        taints: [{ key: "kubernetes.taint", value: "true:NoSchedule" }],
        upgradeSettings: { maxSurge: "1" },
        vmSkuName: "NC_P46_224_v1",
      },
    ],
    kubernetesVersion: "1.XX.Y",
    location: "location",
    managedResourceGroupConfiguration: {
      name: "my-managed-rg",
      location: "East US",
    },
    networkConfiguration: {
      attachedNetworkConfiguration: {
        l2Networks: [
          {
            networkId:
              "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l2Networks/l2NetworkName",
            pluginType: "DPDK",
          },
        ],
        l3Networks: [
          {
            ipamEnabled: "False",
            networkId:
              "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName",
            pluginType: "SRIOV",
          },
        ],
        trunkedNetworks: [
          {
            networkId:
              "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/trunkedNetworks/trunkedNetworkName",
            pluginType: "MACVLAN",
          },
        ],
      },
      cloudServicesNetworkId:
        "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/cloudServicesNetworks/cloudServicesNetworkName",
      cniNetworkId:
        "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName",
      dnsServiceIp: "198.51.101.2",
      l2ServiceLoadBalancerConfiguration: {
        ipAddressPools: [
          {
            name: "pool1",
            addresses: ["198.51.102.2-198.51.102.254"],
            autoAssign: "True",
          },
        ],
      },
      podCidrs: ["198.51.100.0/24"],
      serviceCidrs: ["198.51.101.0/24"],
    },
    tags: { key1: "myvalue1", key2: "myvalue2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkCloud(credential, subscriptionId);
  const result = await client.kubernetesClusters.beginCreateOrUpdateAndWait(
    resourceGroupName,
    kubernetesClusterName,
    kubernetesClusterParameters,
  );
  console.log(result);
}
