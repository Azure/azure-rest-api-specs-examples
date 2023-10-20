/** Samples for MarketplaceGalleryImagesOperation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/ListMarketplaceGalleryImageBySubscription.json
     */
    /**
     * Sample code: ListMarketplaceGalleryImageBySubscription.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void listMarketplaceGalleryImageBySubscription(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.marketplaceGalleryImagesOperations().list(com.azure.core.util.Context.NONE);
    }
}
