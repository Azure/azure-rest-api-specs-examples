/** Samples for AttachedDatabaseConfigurations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoAttachedDatabaseConfigurationsGet.json
     */
    /**
     * Sample code: AttachedDatabaseConfigurationsGet.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void attachedDatabaseConfigurationsGet(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .attachedDatabaseConfigurations()
            .getWithResponse(
                "kustorptest", "kustoCluster2", "attachedDatabaseConfigurationsTest", com.azure.core.util.Context.NONE);
    }
}
