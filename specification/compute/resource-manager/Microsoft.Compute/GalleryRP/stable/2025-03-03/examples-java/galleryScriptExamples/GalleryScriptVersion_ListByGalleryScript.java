
/**
 * Samples for GalleryScriptVersions ListByGalleryScript.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/
     * galleryScriptExamples/GalleryScriptVersion_ListByGalleryScript.json
     */
    /**
     * Sample code: List gallery Script Versions in a gallery Script Definition.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listGalleryScriptVersionsInAGalleryScriptDefinition(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryScriptVersions().listByGalleryScript(
            "myResourceGroupName", "myGalleryName", "myGalleryScriptName", com.azure.core.util.Context.NONE);
    }
}
