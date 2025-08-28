
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/search/resource-manager/Microsoft.Search/stable/2025-05-01/examples/SearchListOperations.json
     */
    /**
     * Sample code: SearchListOperations.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void searchListOperations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.searchServices().manager().serviceClient().getOperations().list(com.azure.core.util.Context.NONE);
    }
}
