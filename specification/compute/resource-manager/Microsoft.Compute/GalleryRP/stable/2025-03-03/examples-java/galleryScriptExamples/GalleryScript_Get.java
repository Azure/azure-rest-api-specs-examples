
/**
 * Samples for GalleryScripts Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/
     * galleryScriptExamples/GalleryScript_Get.json
     */
    /**
     * Sample code: Get a gallery Script.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAGalleryScript(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryScripts().getWithResponse("myResourceGroup",
            "myGalleryName", "myGalleryScriptName", com.azure.core.util.Context.NONE);
    }
}
