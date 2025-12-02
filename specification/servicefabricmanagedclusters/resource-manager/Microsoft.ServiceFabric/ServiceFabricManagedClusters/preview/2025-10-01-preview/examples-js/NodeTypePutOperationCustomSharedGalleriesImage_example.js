const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Service Fabric node type of a given managed cluster.
 *
 * @summary create or update a Service Fabric node type of a given managed cluster.
 * x-ms-original-file: 2025-10-01-preview/NodeTypePutOperationCustomSharedGalleriesImage_example.json
 */
async function putNodeTypeWithSharedGalleriesCustomVmImage() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const result = await client.nodeTypes.createOrUpdate("resRg", "myCluster", "BE", {
    properties: {
      dataDiskSizeGB: 200,
      isPrimary: false,
      vmInstanceCount: 10,
      vmSharedGalleryImageId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-custom-image/providers/Microsoft.Compute/sharedGalleries/35349201-a0b3-405e-8a23-9f1450984307-SFSHAREDGALLERY/images/TestNoProdContainerDImage/versions/latest",
      vmSize: "Standard_D3",
    },
  });
  console.log(result);
}
