
/**
 * Samples for GalleryInVMAccessControlProfiles Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/galleryResourceProfileExamples/GalleryInVMAccessControlProfile_Delete.json
     */
    /**
     * Sample code: Delete a gallery inVMAccessControlProfile.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        deleteAGalleryInVMAccessControlProfile(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryInVMAccessControlProfiles().delete("myResourceGroup", "myGalleryName",
            "myInVMAccessControlProfileName", com.azure.core.util.Context.NONE);
    }
}
