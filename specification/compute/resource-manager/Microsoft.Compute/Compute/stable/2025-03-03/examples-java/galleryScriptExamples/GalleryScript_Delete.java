
/**
 * Samples for GalleryScripts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/galleryScriptExamples/GalleryScript_Delete.json
     */
    /**
     * Sample code: Delete a gallery Script.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void deleteAGalleryScript(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryScripts().delete("myResourceGroup", "myGalleryName", "myGalleryScriptName",
            com.azure.core.util.Context.NONE);
    }
}
