/** Samples for AttachedDatabaseConfigurations ListByCluster. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoAttachedDatabaseConfigurationsListByCluster.json
     */
    /**
     * Sample code: KustoAttachedDatabaseConfigurationsListByCluster.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoAttachedDatabaseConfigurationsListByCluster(
        com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .attachedDatabaseConfigurations()
            .listByCluster("kustorptest", "kustoCluster2", com.azure.core.util.Context.NONE);
    }
}
