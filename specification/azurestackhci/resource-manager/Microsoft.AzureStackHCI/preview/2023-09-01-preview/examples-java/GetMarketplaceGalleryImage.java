/** Samples for MarketplaceGalleryImagesOperation GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/GetMarketplaceGalleryImage.json
     */
    /**
     * Sample code: GetMarketplaceGalleryImage.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void getMarketplaceGalleryImage(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager
            .marketplaceGalleryImagesOperations()
            .getByResourceGroupWithResponse(
                "test-rg", "test-marketplace-gallery-image", com.azure.core.util.Context.NONE);
    }
}
