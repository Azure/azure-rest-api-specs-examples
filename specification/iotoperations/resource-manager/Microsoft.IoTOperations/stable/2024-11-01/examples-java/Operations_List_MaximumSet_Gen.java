
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/Operations_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void operationsList(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
