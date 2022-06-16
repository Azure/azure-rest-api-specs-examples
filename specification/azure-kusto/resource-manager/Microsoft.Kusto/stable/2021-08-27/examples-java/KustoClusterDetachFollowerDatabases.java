import com.azure.core.util.Context;
import com.azure.resourcemanager.kusto.fluent.models.FollowerDatabaseDefinitionInner;

/** Samples for Clusters DetachFollowerDatabases. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoClusterDetachFollowerDatabases.json
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
                "kustoclusterrptest4",
                new FollowerDatabaseDefinitionInner()
                    .withClusterResourceId(
                        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/clusters/leader4")
                    .withAttachedDatabaseConfigurationName("myAttachedDatabaseConfiguration"),
                Context.NONE);
    }
}
