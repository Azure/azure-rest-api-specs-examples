/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/help/resource-manager/Microsoft.Help/preview/2023-09-01-preview/examples/ListOperations.json
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
