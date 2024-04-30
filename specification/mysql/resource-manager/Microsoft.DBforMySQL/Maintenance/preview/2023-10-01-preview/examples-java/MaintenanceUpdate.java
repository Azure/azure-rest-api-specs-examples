
import com.azure.resourcemanager.mysqlflexibleserver.models.MaintenanceUpdate;
import java.time.OffsetDateTime;

/**
 * Samples for Maintenances Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/Maintenance/preview/2023-10-01-preview/examples/
     * MaintenanceUpdate.json
     */
    /**
     * Sample code: Update maintenance on a server.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void updateMaintenanceOnAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.maintenances().update("TestGroup", "testserver", "_T9Q-TS8",
            new MaintenanceUpdate().withMaintenanceStartTime(OffsetDateTime.parse("2024-01-20T00:00:00")),
            com.azure.core.util.Context.NONE);
    }
}
