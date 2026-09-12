
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/RelayOperations_List.json
     */
    /**
     * Sample code: RelayOperationsList.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void relayOperationsList(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
