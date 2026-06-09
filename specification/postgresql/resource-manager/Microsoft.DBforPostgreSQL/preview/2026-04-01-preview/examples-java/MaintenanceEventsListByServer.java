
/**
 * Samples for MaintenanceEvents List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/MaintenanceEventsListByServer.json
     */
    /**
     * Sample code: List ongoing and scheduled maintenance events for a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void listOngoingAndScheduledMaintenanceEventsForAServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.maintenanceEvents().list("exampleresourcegroup", "exampleserver", null,
            com.azure.core.util.Context.NONE);
    }
}
