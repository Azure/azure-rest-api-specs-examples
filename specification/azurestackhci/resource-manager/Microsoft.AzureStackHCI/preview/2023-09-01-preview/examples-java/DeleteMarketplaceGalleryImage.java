/** Samples for MarketplaceGalleryImagesOperation Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/DeleteMarketplaceGalleryImage.json
     */
    /**
     * Sample code: DeleteMarketplaceGalleryImage.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void deleteMarketplaceGalleryImage(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager
            .marketplaceGalleryImagesOperations()
            .delete("test-rg", "test-marketplace-gallery-image", com.azure.core.util.Context.NONE);
    }
}
