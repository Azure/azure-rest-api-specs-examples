import com.azure.core.util.Context;

/** Samples for Scripts Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoScriptsGet.json
     */
    /**
     * Sample code: KustoScriptsGet.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoScriptsGet(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .scripts()
            .getWithResponse("kustorptest", "kustoclusterrptest4", "Kustodatabase8", "kustoScript1", Context.NONE);
    }
}
