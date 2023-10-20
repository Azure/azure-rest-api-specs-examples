/** Samples for GalleryImagesOperation GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/GetGalleryImage.json
     */
    /**
     * Sample code: GetGalleryImage.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void getGalleryImage(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager
            .galleryImagesOperations()
            .getByResourceGroupWithResponse("test-rg", "test-gallery-image", com.azure.core.util.Context.NONE);
    }
}
