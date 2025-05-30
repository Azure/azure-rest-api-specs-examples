
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-01-preview/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void operationsList(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
