
/**
 * Samples for GalleryInVMAccessControlProfileVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/
     * galleryResourceProfileExamples/GalleryInVMAccessControlProfileVersion_Delete.json
     */
    /**
     * Sample code: Delete a gallery inVMAccessControlProfile version.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        deleteAGalleryInVMAccessControlProfileVersion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryInVMAccessControlProfileVersions().delete(
            "myResourceGroup", "myGalleryName", "myInVMAccessControlProfileName", "1.0.0",
            com.azure.core.util.Context.NONE);
    }
}
