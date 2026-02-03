
/**
 * Samples for GalleryInVMAccessControlProfileVersions ListByGalleryInVMAccessControlProfile.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/
     * galleryResourceProfileExamples/GalleryInVMAccessControlProfileVersion_ListByGalleryInVMAccessControlProfile.json
     */
    /**
     * Sample code: List gallery inVMAccessControlProfile versions in a gallery inVMAccessControlProfile.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listGalleryInVMAccessControlProfileVersionsInAGalleryInVMAccessControlProfile(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryInVMAccessControlProfileVersions()
            .listByGalleryInVMAccessControlProfile("myResourceGroup", "myGalleryName", "myInVMAccessControlProfileName",
                com.azure.core.util.Context.NONE);
    }
}
