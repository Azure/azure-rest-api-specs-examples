
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/Operations_List.json
     */
    /**
     * Sample code: Lists all Chaos Studio operations.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void listsAllChaosStudioOperations(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
