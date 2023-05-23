/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/OperationsGet.json
     */
    /**
     * Sample code: OperationsGet.
     *
     * @param manager Entry point to DataBoxManager.
     */
    public static void operationsGet(com.azure.resourcemanager.databox.DataBoxManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
