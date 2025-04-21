
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to StandbyPoolManager.
     */
    public static void operationsList(com.azure.resourcemanager.standbypool.StandbyPoolManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
