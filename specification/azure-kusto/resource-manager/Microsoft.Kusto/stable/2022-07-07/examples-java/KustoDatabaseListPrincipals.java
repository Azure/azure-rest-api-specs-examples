import com.azure.core.util.Context;

/** Samples for Databases ListPrincipals. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-07-07/examples/KustoDatabaseListPrincipals.json
     */
    /**
     * Sample code: KustoDatabaseListPrincipals.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDatabaseListPrincipals(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.databases().listPrincipals("kustorptest", "kustoCluster", "KustoDatabase8", Context.NONE);
    }
}
