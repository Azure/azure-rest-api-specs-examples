
import com.azure.resourcemanager.compute.fluent.models.SharingUpdateInner;
import com.azure.resourcemanager.compute.models.SharingUpdateOperationTypes;

/**
 * Samples for GallerySharingProfile Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/galleryExamples/Gallery_ResetSharingProfile.json
     */
    /**
     * Sample code: reset sharing profile of a gallery.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void resetSharingProfileOfAGallery(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGallerySharingProfiles().update("myResourceGroup", "myGalleryName",
            new SharingUpdateInner().withOperationType(SharingUpdateOperationTypes.RESET),
            com.azure.core.util.Context.NONE);
    }
}
