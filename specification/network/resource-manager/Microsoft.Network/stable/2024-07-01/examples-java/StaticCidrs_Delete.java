
/**
 * Samples for StaticCidrs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/StaticCidrs_Delete.json
     */
    /**
     * Sample code: StaticCidrs_Delete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void staticCidrsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getStaticCidrs().delete("rg1", "TestNetworkManager", "TestPool",
            "TestStaticCidr", com.azure.core.util.Context.NONE);
    }
}
