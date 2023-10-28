/** Samples for Services List. */
public final class Main {
    /*
     * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2023-11-01/examples/SearchListServicesBySubscription.json
     */
    /**
     * Sample code: SearchListServicesBySubscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void searchListServicesBySubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.searchServices().manager().serviceClient().getServices().list(null, com.azure.core.util.Context.NONE);
    }
}
