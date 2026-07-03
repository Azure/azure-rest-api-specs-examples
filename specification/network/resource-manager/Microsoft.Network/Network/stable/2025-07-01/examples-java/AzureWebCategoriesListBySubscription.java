
/**
 * Samples for WebCategories List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/AzureWebCategoriesListBySubscription.json
     */
    /**
     * Sample code: List all Azure Web Categories for a given subscription.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listAllAzureWebCategoriesForAGivenSubscription(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getWebCategories().list(com.azure.core.util.Context.NONE);
    }
}
