
/**
 * Samples for GalleryScriptVersions ListByGalleryScript.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/galleryScriptExamples/GalleryScriptVersion_ListByGalleryScript.json
     */
    /**
     * Sample code: List gallery Script Versions in a gallery Script Definition.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        listGalleryScriptVersionsInAGalleryScriptDefinition(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryScriptVersions().listByGalleryScript("myResourceGroupName", "myGalleryName",
            "myGalleryScriptName", com.azure.core.util.Context.NONE);
    }
}
