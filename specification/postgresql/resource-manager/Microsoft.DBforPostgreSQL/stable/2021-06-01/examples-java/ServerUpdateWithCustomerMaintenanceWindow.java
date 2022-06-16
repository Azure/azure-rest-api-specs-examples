import com.azure.core.util.Context;
import com.azure.resourcemanager.postgresqlflexibleserver.models.CreateModeForUpdate;
import com.azure.resourcemanager.postgresqlflexibleserver.models.MaintenanceWindow;
import com.azure.resourcemanager.postgresqlflexibleserver.models.Server;

/** Samples for Servers Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2021-06-01/examples/ServerUpdateWithCustomerMaintenanceWindow.json
     */
    /**
     * Sample code: ServerUpdateWithCustomerMaintenanceWindow.
     *
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void serverUpdateWithCustomerMaintenanceWindow(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        Server resource =
            manager.servers().getByResourceGroupWithResponse("testrg", "pgtestsvc4", Context.NONE).getValue();
        resource
            .update()
            .withMaintenanceWindow(
                new MaintenanceWindow()
                    .withCustomWindow("Enabled")
                    .withStartHour(8)
                    .withStartMinute(0)
                    .withDayOfWeek(0))
            .withCreateMode(CreateModeForUpdate.UPDATE)
            .apply();
    }
}
