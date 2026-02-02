
import com.azure.resourcemanager.compute.fluent.models.SharingUpdateInner;
import com.azure.resourcemanager.compute.models.SharingUpdateOperationTypes;

/**
 * Samples for GallerySharingProfile Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/
     * Gallery_EnableCommunityGallery.json
     */
    /**
     * Sample code: share a gallery to community.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void shareAGalleryToCommunity(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGallerySharingProfiles().update("myResourceGroup",
            "myGalleryName", new SharingUpdateInner().withOperationType(SharingUpdateOperationTypes.ENABLE_COMMUNITY),
            com.azure.core.util.Context.NONE);
    }
}
