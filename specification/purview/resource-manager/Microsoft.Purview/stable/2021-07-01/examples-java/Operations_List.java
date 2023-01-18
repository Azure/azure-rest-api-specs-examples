/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to PurviewManager.
     */
    public static void operationsList(com.azure.resourcemanager.purview.PurviewManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
