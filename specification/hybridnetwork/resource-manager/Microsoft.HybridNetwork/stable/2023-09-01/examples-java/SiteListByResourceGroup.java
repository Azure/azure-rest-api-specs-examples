
/**
 * Samples for Sites ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * SiteListByResourceGroup.json
     */
    /**
     * Sample code: List all network sites.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void listAllNetworkSites(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.sites().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
