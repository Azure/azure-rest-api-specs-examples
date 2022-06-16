import com.azure.core.util.Context;

/** Samples for Databases Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoDatabasesGet.json
     */
    /**
     * Sample code: KustoDatabasesGet.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDatabasesGet(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.databases().getWithResponse("kustorptest", "kustoclusterrptest4", "KustoDatabase8", Context.NONE);
    }
}
