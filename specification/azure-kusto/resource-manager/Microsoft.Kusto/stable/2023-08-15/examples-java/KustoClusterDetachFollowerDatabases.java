import com.azure.resourcemanager.kusto.fluent.models.FollowerDatabaseDefinitionInner;

/** Samples for Clusters DetachFollowerDatabases. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoClusterDetachFollowerDatabases.json
     */
    /**
     * Sample code: KustoClusterDetachFollowerDatabases.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClusterDetachFollowerDatabases(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .clusters()
            .detachFollowerDatabases(
                "kustorptest",
                "kustoCluster",
                new FollowerDatabaseDefinitionInner()
                    .withClusterResourceId(
                        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/clusters/kustoCluster2")
                    .withAttachedDatabaseConfigurationName("attachedDatabaseConfigurationsTest"),
                com.azure.core.util.Context.NONE);
    }
}
