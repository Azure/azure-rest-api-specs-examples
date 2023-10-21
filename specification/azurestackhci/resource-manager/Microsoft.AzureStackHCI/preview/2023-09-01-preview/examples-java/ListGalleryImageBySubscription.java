/** Samples for GalleryImagesOperation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/ListGalleryImageBySubscription.json
     */
    /**
     * Sample code: ListGalleryImageBySubscription.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void listGalleryImageBySubscription(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.galleryImagesOperations().list(com.azure.core.util.Context.NONE);
    }
}
