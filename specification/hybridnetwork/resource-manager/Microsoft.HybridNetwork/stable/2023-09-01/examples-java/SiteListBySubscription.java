
/**
 * Samples for Sites List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * SiteListBySubscription.json
     */
    /**
     * Sample code: List all hybrid network sites in a subscription.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void
        listAllHybridNetworkSitesInASubscription(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.sites().list(com.azure.core.util.Context.NONE);
    }
}
