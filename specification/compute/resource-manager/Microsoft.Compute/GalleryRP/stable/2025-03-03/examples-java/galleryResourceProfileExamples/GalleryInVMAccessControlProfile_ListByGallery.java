
/**
 * Samples for GalleryInVMAccessControlProfiles ListByGallery.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/
     * galleryResourceProfileExamples/GalleryInVMAccessControlProfile_ListByGallery.json
     */
    /**
     * Sample code: List gallery inVMAccessControlProfiles in a gallery.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listGalleryInVMAccessControlProfilesInAGallery(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryInVMAccessControlProfiles()
            .listByGallery("myResourceGroup", "myGalleryName", com.azure.core.util.Context.NONE);
    }
}
