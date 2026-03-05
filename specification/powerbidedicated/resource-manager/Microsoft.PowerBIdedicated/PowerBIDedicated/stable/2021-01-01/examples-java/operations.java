
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2021-01-01/operations.json
     */
    /**
     * Sample code: List operations.
     * 
     * @param manager Entry point to PowerBIDedicatedManager.
     */
    public static void listOperations(com.azure.resourcemanager.powerbidedicated.PowerBIDedicatedManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
