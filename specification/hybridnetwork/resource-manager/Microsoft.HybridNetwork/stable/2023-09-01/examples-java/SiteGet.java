
/**
 * Samples for Sites GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/SiteGet.json
     */
    /**
     * Sample code: Get network site.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getNetworkSite(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.sites().getByResourceGroupWithResponse("rg1", "testSite", com.azure.core.util.Context.NONE);
    }
}
