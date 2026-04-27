
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01-preview/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to AppnetworkManager.
     */
    public static void operationsList(com.azure.resourcemanager.appnetwork.AppnetworkManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
