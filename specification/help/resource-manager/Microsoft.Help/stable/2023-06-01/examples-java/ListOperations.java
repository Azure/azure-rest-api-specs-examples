/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/help/resource-manager/Microsoft.Help/stable/2023-06-01/examples/ListOperations.json
     */
    /**
     * Sample code: List All Operations.
     *
     * @param manager Entry point to SelfHelpManager.
     */
    public static void listAllOperations(com.azure.resourcemanager.selfhelp.SelfHelpManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
