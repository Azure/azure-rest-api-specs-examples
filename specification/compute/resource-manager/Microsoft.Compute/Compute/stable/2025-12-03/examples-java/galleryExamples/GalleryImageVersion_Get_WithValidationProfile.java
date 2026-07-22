
import com.azure.resourcemanager.compute.models.ReplicationStatusTypes;

/**
 * Samples for GalleryImageVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/galleryExamples/GalleryImageVersion_Get_WithValidationProfile.json
     */
    /**
     * Sample code: Get a gallery image version with validation profile.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getAGalleryImageVersionWithValidationProfile(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryImageVersions().getWithResponse("myResourceGroup", "myGalleryName",
            "myGalleryImageName", "1.0.0", ReplicationStatusTypes.fromString("ValidationProfile"),
            com.azure.core.util.Context.NONE);
    }
}
