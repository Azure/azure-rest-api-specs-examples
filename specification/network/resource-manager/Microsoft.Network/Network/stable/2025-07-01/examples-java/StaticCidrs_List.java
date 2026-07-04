
/**
 * Samples for StaticCidrs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/StaticCidrs_List.json
     */
    /**
     * Sample code: StaticCidrs_List.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void staticCidrsList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getStaticCidrs().list("rg1", "TestNetworkManager", "TestPool", null, null, null, null,
            null, com.azure.core.util.Context.NONE);
    }
}
