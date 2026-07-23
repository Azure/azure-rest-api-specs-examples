
/**
 * Samples for GalleryInVMAccessControlProfileVersions ListByGalleryInVMAccessControlProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/galleryResourceProfileExamples/
     * GalleryInVMAccessControlProfileVersion_ListByGalleryInVMAccessControlProfile.json
     */
    /**
     * Sample code: List gallery inVMAccessControlProfile versions in a gallery inVMAccessControlProfile.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listGalleryInVMAccessControlProfileVersionsInAGalleryInVMAccessControlProfile(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryInVMAccessControlProfileVersions().listByGalleryInVMAccessControlProfile(
            "myResourceGroup", "myGalleryName", "myInVMAccessControlProfileName", com.azure.core.util.Context.NONE);
    }
}
