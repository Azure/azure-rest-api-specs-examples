
/**
 * Samples for GalleryApplications Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/galleryExamples/GalleryApplication_Delete.json
     */
    /**
     * Sample code: Delete a gallery Application.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void deleteAGalleryApplication(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryApplications().delete("myResourceGroup", "myGalleryName",
            "myGalleryApplicationName", com.azure.core.util.Context.NONE);
    }
}
