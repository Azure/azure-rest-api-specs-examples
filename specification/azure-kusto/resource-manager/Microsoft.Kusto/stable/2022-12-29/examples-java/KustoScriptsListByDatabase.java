/** Samples for Scripts ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-12-29/examples/KustoScriptsListByDatabase.json
     */
    /**
     * Sample code: KustoScriptsList.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoScriptsList(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .scripts()
            .listByDatabase("kustorptest", "kustoCluster", "Kustodatabase8", com.azure.core.util.Context.NONE);
    }
}
