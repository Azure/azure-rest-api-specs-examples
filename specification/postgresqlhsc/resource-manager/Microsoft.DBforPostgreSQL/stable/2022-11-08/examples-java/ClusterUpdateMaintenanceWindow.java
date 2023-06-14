import com.azure.resourcemanager.cosmosdbforpostgresql.models.Cluster;
import com.azure.resourcemanager.cosmosdbforpostgresql.models.MaintenanceWindow;

/** Samples for Clusters Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresqlhsc/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-11-08/examples/ClusterUpdateMaintenanceWindow.json
     */
    /**
     * Sample code: Update or define maintenance window.
     *
     * @param manager Entry point to CosmosDBForPostgreSqlManager.
     */
    public static void updateOrDefineMaintenanceWindow(
        com.azure.resourcemanager.cosmosdbforpostgresql.CosmosDBForPostgreSqlManager manager) {
        Cluster resource =
            manager
                .clusters()
                .getByResourceGroupWithResponse("TestGroup", "testcluster", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withMaintenanceWindow(
                new MaintenanceWindow()
                    .withCustomWindow("Enabled")
                    .withStartHour(8)
                    .withStartMinute(0)
                    .withDayOfWeek(0))
            .apply();
    }
}
