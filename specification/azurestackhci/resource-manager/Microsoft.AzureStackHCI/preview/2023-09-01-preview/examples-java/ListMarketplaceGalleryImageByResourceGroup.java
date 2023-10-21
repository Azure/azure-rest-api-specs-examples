/** Samples for MarketplaceGalleryImagesOperation ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/ListMarketplaceGalleryImageByResourceGroup.json
     */
    /**
     * Sample code: ListMarketplaceGalleryImageByResourceGroup.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void listMarketplaceGalleryImageByResourceGroup(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.marketplaceGalleryImagesOperations().listByResourceGroup("test-rg", com.azure.core.util.Context.NONE);
    }
}
