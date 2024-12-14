
/**
 * Samples for NetworkFunctions ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * NetworkFunctionListByResourceGroup.json
     */
    /**
     * Sample code: List network function in resource group.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void
        listNetworkFunctionInResourceGroup(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkFunctions().listByResourceGroup("rg", com.azure.core.util.Context.NONE);
    }
}
