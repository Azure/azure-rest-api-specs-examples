
/**
 * Samples for GalleryScripts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/
     * galleryScriptExamples/GalleryScript_Delete.json
     */
    /**
     * Sample code: Delete a gallery Script.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAGalleryScript(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryScripts().delete("myResourceGroup", "myGalleryName",
            "myGalleryScriptName", com.azure.core.util.Context.NONE);
    }
}
