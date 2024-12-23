
/**
 * Samples for StaticCidrs List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/StaticCidrs_List.json
     */
    /**
     * Sample code: StaticCidrs_List.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void staticCidrsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getStaticCidrs().list("rg1", "TestNetworkManager", "TestPool", null,
            null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
