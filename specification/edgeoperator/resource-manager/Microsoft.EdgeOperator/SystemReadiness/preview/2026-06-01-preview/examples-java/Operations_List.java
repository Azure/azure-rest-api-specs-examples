
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to EdgeOperatorManager.
     */
    public static void operationsList(com.azure.resourcemanager.edgeoperator.EdgeOperatorManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
