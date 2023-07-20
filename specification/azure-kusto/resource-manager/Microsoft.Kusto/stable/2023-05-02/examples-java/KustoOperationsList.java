/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoOperationsList.json
     */
    /**
     * Sample code: KustoOperationsList.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoOperationsList(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
