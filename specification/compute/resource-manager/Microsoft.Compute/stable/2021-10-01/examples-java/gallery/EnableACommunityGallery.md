Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.SharingUpdateInner;
import com.azure.resourcemanager.compute.models.SharingUpdateOperationTypes;

/** Samples for GallerySharingProfile Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/EnableACommunityGallery.json
     */
    /**
     * Sample code: share a gallery to community.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void shareAGalleryToCommunity(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGallerySharingProfiles()
            .update(
                "myResourceGroup",
                "myGalleryName",
                new SharingUpdateInner().withOperationType(SharingUpdateOperationTypes.ENABLE_COMMUNITY),
                Context.NONE);
    }
}
```
