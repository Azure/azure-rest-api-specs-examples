
/**
 * Samples for CommunityGalleryImages List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/
     * communityGalleryExamples/CommunityGalleryImage_List.json
     */
    /**
     * Sample code: List community gallery images.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listCommunityGalleryImages(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCommunityGalleryImages().list("myLocation",
            "publicGalleryName", com.azure.core.util.Context.NONE);
    }
}
