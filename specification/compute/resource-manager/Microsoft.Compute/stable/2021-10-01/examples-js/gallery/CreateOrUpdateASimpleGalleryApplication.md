Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_18.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a gallery Application Definition.
 *
 * @summary Create or update a gallery Application Definition.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/CreateOrUpdateASimpleGalleryApplication.json
 */
async function createOrUpdateASimpleGalleryApplication() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const galleryName = "myGalleryName";
  const galleryApplicationName = "myGalleryApplicationName";
  const galleryApplication = {
    description: "This is the gallery application description.",
    eula: "This is the gallery application EULA.",
    location: "West US",
    privacyStatementUri: "myPrivacyStatementUri}",
    releaseNoteUri: "myReleaseNoteUri",
    supportedOSType: "Windows",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.galleryApplications.beginCreateOrUpdateAndWait(
    resourceGroupName,
    galleryName,
    galleryApplicationName,
    galleryApplication
  );
  console.log(result);
}

createOrUpdateASimpleGalleryApplication().catch(console.error);
```
