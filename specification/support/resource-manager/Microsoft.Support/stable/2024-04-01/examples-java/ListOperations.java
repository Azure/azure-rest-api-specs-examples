
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/support/resource-manager/Microsoft.Support/stable/2024-04-01/examples/ListOperations.json
     */
    /**
     * Sample code: Get all operations.
     * 
     * @param manager Entry point to SupportManager.
     */
    public static void getAllOperations(com.azure.resourcemanager.support.SupportManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
