
/**
 * Samples for GalleryInVMAccessControlProfiles Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2024-03-03/examples/
     * galleryResourceProfileExamples/GalleryInVMAccessControlProfile_Get.json
     */
    /**
     * Sample code: Get a gallery inVMAccessControlProfile.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAGalleryInVMAccessControlProfile(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryInVMAccessControlProfiles().getWithResponse(
            "myResourceGroup", "myGalleryName", "myInVMAccessControlProfileName", com.azure.core.util.Context.NONE);
    }
}
