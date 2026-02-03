
/**
 * Samples for GalleryInVMAccessControlProfiles Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/
     * galleryResourceProfileExamples/GalleryInVMAccessControlProfile_Delete.json
     */
    /**
     * Sample code: Delete a gallery inVMAccessControlProfile.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAGalleryInVMAccessControlProfile(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryInVMAccessControlProfiles().delete(
            "myResourceGroup", "myGalleryName", "myInVMAccessControlProfileName", com.azure.core.util.Context.NONE);
    }
}
