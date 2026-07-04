
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/OperationList.json
     */
    /**
     * Sample code: Get a list of operations for a resource provider.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getAListOfOperationsForAResourceProvider(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getOperations().list(com.azure.core.util.Context.NONE);
    }
}
