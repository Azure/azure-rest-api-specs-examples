/** Samples for GalleryImagesOperation ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/ListGalleryImageByResourceGroup.json
     */
    /**
     * Sample code: ListGalleryImageByResourceGroup.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void listGalleryImageByResourceGroup(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.galleryImagesOperations().listByResourceGroup("test-rg", com.azure.core.util.Context.NONE);
    }
}
