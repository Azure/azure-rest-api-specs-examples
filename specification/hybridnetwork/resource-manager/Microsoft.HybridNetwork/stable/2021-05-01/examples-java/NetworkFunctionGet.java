import com.azure.core.util.Context;

/** Samples for NetworkFunctions GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/NetworkFunctionGet.json
     */
    /**
     * Sample code: Get network function resource.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getNetworkFunctionResource(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkFunctions().getByResourceGroupWithResponse("rg", "testNf", Context.NONE);
    }
}
