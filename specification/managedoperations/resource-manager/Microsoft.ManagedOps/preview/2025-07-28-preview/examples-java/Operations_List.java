
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-28-preview/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to ManagedOpsManager.
     */
    public static void operationsList(com.azure.resourcemanager.managedops.ManagedOpsManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
