
/**
 * Samples for StaticCidrs Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/StaticCidrs_Create.json
     */
    /**
     * Sample code: StaticCidrs_Create.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void staticCidrsCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getStaticCidrs().createWithResponse("rg1", "TestNetworkManager",
            "TestPool", "TestStaticCidr", null, com.azure.core.util.Context.NONE);
    }
}
