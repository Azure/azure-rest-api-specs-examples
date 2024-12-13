
/**
 * Samples for Components ListByNetworkFunction.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * ComponentListByNetworkFunction.json
     */
    /**
     * Sample code: List components in network function.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void
        listComponentsInNetworkFunction(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.components().listByNetworkFunction("rg", "testNf", com.azure.core.util.Context.NONE);
    }
}
