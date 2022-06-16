import com.azure.core.util.Context;

/** Samples for AttachedDatabaseConfigurations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoAttachedDatabaseConfigurationsDelete.json
     */
    /**
     * Sample code: AttachedDatabaseConfigurationsDelete.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void attachedDatabaseConfigurationsDelete(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .attachedDatabaseConfigurations()
            .delete("kustorptest", "kustoCluster", "attachedDatabaseConfigurationsTest", Context.NONE);
    }
}
