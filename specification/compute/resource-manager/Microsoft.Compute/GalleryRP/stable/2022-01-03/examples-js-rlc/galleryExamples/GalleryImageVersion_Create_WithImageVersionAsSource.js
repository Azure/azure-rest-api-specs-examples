const createComputeManagementClient = require("@azure-rest/arm-compute").default,
  { getLongRunningPoller } = require("@azure-rest/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Create or update a gallery image version.
 *
 * @summary Create or update a gallery image version.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-01-03/examples/galleryExamples/GalleryImageVersion_Create_WithImageVersionAsSource.json
 */
async function createOrUpdateASimpleGalleryImageVersionUsingSharedImageAsSource() {
  const credential = new DefaultAzureCredential();
  const client = createComputeManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "myResourceGroup";
  const galleryName = "myGalleryName";
  const galleryImageName = "myGalleryImageName";
  const galleryImageVersionName = "1.0.0";
  const options = {
    body: {
      location: "West US",
      properties: {
        publishingProfile: {
          targetRegions: [
            {
              name: "West US",
              encryption: {
                dataDiskImages: [
                  {
                    diskEncryptionSetId:
                      "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherWestUSDiskEncryptionSet",
                    lun: 0,
                  },
                  {
                    diskEncryptionSetId:
                      "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet",
                    lun: 1,
                  },
                ],
                osDiskImage: {
                  diskEncryptionSetId:
                    "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet",
                },
              },
              regionalReplicaCount: 1,
            },
            {
              name: "East US",
              encryption: {
                dataDiskImages: [
                  {
                    diskEncryptionSetId:
                      "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherEastUSDiskEncryptionSet",
                    lun: 0,
                  },
                  {
                    diskEncryptionSetId:
                      "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet",
                    lun: 1,
                  },
                ],
                osDiskImage: {
                  diskEncryptionSetId:
                    "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet",
                },
              },
              regionalReplicaCount: 2,
              storageAccountType: "Standard_ZRS",
            },
          ],
        },
        storageProfile: {
          source: {
            id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/galleries/{galleryName}/images/{imageDefinitionName}/versions/{versionName}",
          },
        },
      },
    },
    queryParameters: { "api-version": "2022-01-03" },
  };
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}/versions/{galleryImageVersionName}",
      subscriptionId,
      resourceGroupName,
      galleryName,
      galleryImageName,
      galleryImageVersionName
    )
    .put(options);
  const poller = getLongRunningPoller(client, initialResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

createOrUpdateASimpleGalleryImageVersionUsingSharedImageAsSource().catch(console.error);
