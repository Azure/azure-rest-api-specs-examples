
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/GetOperations.
     * json
     */
    /**
     * Sample code: Get Registration Operations.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void getRegistrationOperations(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
