
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/SBOperations_List.json
     */
    /**
     * Sample code: OperationsList.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void operationsList(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getOperations().list(com.azure.core.util.Context.NONE);
    }
}
