
import com.azure.resourcemanager.postgresqlflexibleserver.models.MaintenanceEventStatusFilter;

/**
 * Samples for MaintenanceEvents List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/MaintenanceEventsListByServerWithFilter.json
     */
    /**
     * Sample code: List maintenance events filtered by status for a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listMaintenanceEventsFilteredByStatusForAServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.maintenanceEvents().list("exampleresourcegroup", "exampleserver", MaintenanceEventStatusFilter.UPCOMING,
            com.azure.core.util.Context.NONE);
    }
}
