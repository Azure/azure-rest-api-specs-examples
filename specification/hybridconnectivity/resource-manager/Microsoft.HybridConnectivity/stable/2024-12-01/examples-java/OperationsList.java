
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/OperationsList.json
     */
    /**
     * Sample code: HybridConnectivityOperationsList.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void hybridConnectivityOperationsList(
        com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
