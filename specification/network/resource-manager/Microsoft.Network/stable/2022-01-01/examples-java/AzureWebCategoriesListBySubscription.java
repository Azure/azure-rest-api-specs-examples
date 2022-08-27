import com.azure.core.util.Context;

/** Samples for WebCategories List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/AzureWebCategoriesListBySubscription.json
     */
    /**
     * Sample code: List all Azure Web Categories for a given subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllAzureWebCategoriesForAGivenSubscription(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getWebCategories().list(Context.NONE);
    }
}
