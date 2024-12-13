
/**
 * Samples for Sites Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/SiteDelete.json
     */
    /**
     * Sample code: Delete network site.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void deleteNetworkSite(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.sites().delete("rg1", "testSite", com.azure.core.util.Context.NONE);
    }
}
