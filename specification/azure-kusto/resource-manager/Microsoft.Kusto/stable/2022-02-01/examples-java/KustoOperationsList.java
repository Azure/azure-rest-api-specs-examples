import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoOperationsList.json
     */
    /**
     * Sample code: KustoOperationsList.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoOperationsList(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.operations().list(Context.NONE);
    }
}
