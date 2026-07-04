
/**
 * Samples for StaticCidrs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/StaticCidrs_Delete.json
     */
    /**
     * Sample code: StaticCidrs_Delete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void staticCidrsDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getStaticCidrs().delete("rg1", "TestNetworkManager", "TestPool", "TestStaticCidr",
            com.azure.core.util.Context.NONE);
    }
}
