
/**
 * Samples for GalleryInVMAccessControlProfiles Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/galleryResourceProfileExamples/GalleryInVMAccessControlProfile_Get.json
     */
    /**
     * Sample code: Get a gallery inVMAccessControlProfile.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getAGalleryInVMAccessControlProfile(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryInVMAccessControlProfiles().getWithResponse("myResourceGroup",
            "myGalleryName", "myInVMAccessControlProfileName", com.azure.core.util.Context.NONE);
    }
}
