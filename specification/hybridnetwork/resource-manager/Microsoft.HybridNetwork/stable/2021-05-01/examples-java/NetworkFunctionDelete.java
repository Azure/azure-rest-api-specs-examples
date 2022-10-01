import com.azure.core.util.Context;

/** Samples for NetworkFunctions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2021-05-01/examples/NetworkFunctionDelete.json
     */
    /**
     * Sample code: Delete network function resource.
     *
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void deleteNetworkFunctionResource(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkFunctions().delete("rg", "testNf", Context.NONE);
    }
}
