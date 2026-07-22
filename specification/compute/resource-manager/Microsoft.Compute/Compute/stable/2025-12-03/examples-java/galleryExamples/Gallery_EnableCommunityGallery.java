
import com.azure.resourcemanager.compute.fluent.models.SharingUpdateInner;
import com.azure.resourcemanager.compute.models.SharingUpdateOperationTypes;

/**
 * Samples for GallerySharingProfile Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/galleryExamples/Gallery_EnableCommunityGallery.json
     */
    /**
     * Sample code: share a gallery to community.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void shareAGalleryToCommunity(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGallerySharingProfiles().update("myResourceGroup", "myGalleryName",
            new SharingUpdateInner().withOperationType(SharingUpdateOperationTypes.ENABLE_COMMUNITY),
            com.azure.core.util.Context.NONE);
    }
}
