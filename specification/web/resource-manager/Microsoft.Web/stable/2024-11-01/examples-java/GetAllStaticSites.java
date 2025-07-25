
/**
 * Samples for StaticSites List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/GetAllStaticSites.json
     */
    /**
     * Sample code: Get all static sites in a subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAllStaticSitesInASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getStaticSites().list(com.azure.core.util.Context.NONE);
    }
}
