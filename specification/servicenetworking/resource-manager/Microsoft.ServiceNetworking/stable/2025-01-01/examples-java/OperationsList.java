
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/OperationsList.json
     */
    /**
     * Sample code: Get Operations List.
     * 
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void getOperationsList(com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
