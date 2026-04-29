
/**
 * Samples for GalleryInVMAccessControlProfileVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/galleryResourceProfileExamples/GalleryInVMAccessControlProfileVersion_Delete.json
     */
    /**
     * Sample code: Delete a gallery inVMAccessControlProfile version.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        deleteAGalleryInVMAccessControlProfileVersion(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryInVMAccessControlProfileVersions().delete("myResourceGroup", "myGalleryName",
            "myInVMAccessControlProfileName", "1.0.0", com.azure.core.util.Context.NONE);
    }
}
