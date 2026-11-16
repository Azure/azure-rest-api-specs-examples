
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/EHOperations_List.json
     */
    /**
     * Sample code: EHOperations_List.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void eHOperationsList(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getOperations().list(com.azure.core.util.Context.NONE);
    }
}
