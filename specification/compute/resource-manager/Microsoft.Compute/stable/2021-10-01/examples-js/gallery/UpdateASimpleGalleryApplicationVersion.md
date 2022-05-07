Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_17.3.1/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function updateASimpleGalleryApplicationVersion() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const galleryName = "myGalleryName";
  const galleryApplicationName = "myGalleryApplicationName";
  const galleryApplicationVersionName = "1.0.0";
  const galleryApplicationVersion = {
    publishingProfile: {
      endOfLifeDate: new Date("2019-07-01T07:00:00Z"),
      manageActions: {
        install:
          'powershell -command "Expand-Archive -Path package.zip -DestinationPath C:package"',
        remove: "del C:package ",
      },
      replicaCount: 1,
      source: {
        mediaLink:
          "https://mystorageaccount.blob.core.windows.net/mycontainer/package.zip?{sasKey}",
      },
      storageAccountType: "Standard_LRS",
      targetRegions: [
        {
          name: "West US",
          regionalReplicaCount: 1,
          storageAccountType: "Standard_LRS",
        },
      ],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.galleryApplicationVersions.beginUpdateAndWait(
    resourceGroupName,
    galleryName,
    galleryApplicationName,
    galleryApplicationVersionName,
    galleryApplicationVersion
  );
  console.log(result);
}

updateASimpleGalleryApplicationVersion().catch(console.error);
```
