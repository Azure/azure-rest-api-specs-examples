
/**
 * Samples for SiteNetworkServices Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * SiteNetworkServiceDelete.json
     */
    /**
     * Sample code: Delete network site.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void deleteNetworkSite(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.siteNetworkServices().delete("rg1", "testSiteNetworkServiceName", com.azure.core.util.Context.NONE);
    }
}
