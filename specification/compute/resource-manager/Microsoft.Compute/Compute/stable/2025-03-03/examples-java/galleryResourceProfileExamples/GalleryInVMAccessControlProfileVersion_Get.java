
/**
 * Samples for GalleryInVMAccessControlProfileVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/galleryResourceProfileExamples/GalleryInVMAccessControlProfileVersion_Get.json
     */
    /**
     * Sample code: Get a gallery inVMAccessControlProfile version.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getAGalleryInVMAccessControlProfileVersion(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryInVMAccessControlProfileVersions().getWithResponse("myResourceGroup",
            "myGalleryName", "myInVMAccessControlProfileName", "1.0.0", com.azure.core.util.Context.NONE);
    }
}
