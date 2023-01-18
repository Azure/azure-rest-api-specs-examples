/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to LogzManager.
     */
    public static void operationsList(com.azure.resourcemanager.logz.LogzManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
