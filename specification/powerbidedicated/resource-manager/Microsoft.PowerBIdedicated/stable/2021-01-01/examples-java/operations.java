/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/operations.json
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
