
/**
 * Samples for GalleryInVMAccessControlProfiles ListByGallery.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/galleryResourceProfileExamples/GalleryInVMAccessControlProfile_ListByGallery.json
     */
    /**
     * Sample code: List gallery inVMAccessControlProfiles in a gallery.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        listGalleryInVMAccessControlProfilesInAGallery(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryInVMAccessControlProfiles().listByGallery("myResourceGroup", "myGalleryName",
            com.azure.core.util.Context.NONE);
    }
}
