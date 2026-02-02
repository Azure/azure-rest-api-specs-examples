
/**
 * Samples for CommunityGalleryImageVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/
     * communityGalleryExamples/CommunityGalleryImageVersion_List.json
     */
    /**
     * Sample code: List community gallery image versions.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listCommunityGalleryImageVersions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCommunityGalleryImageVersions().list("myLocation",
            "publicGalleryName", "myGalleryImageName", com.azure.core.util.Context.NONE);
    }
}
