
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-11-01/Operations_List.json
     */
    /**
     * Sample code: List operations.
     * 
     * @param manager Entry point to FabricManager.
     */
    public static void listOperations(com.azure.resourcemanager.fabric.FabricManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
