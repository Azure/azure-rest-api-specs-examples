
/**
 * Samples for Usages ListBySubscription.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/search/resource-manager/Microsoft.Search/stable/2023-11-01/examples/GetQuotaUsagesList.json
     */
    /**
     * Sample code: GetQuotaUsagesList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getQuotaUsagesList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.searchServices().manager().serviceClient().getUsages().listBySubscription("westus", null,
            com.azure.core.util.Context.NONE);
    }
}
