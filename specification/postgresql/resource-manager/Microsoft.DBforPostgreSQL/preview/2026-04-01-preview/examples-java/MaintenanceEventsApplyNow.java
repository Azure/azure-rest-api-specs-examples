
/**
 * Samples for MaintenanceEvents ApplyNow.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/MaintenanceEventsApplyNow.json
     */
    /**
     * Sample code: Apply maintenance immediately for a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void applyMaintenanceImmediatelyForAServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.maintenanceEvents().applyNow("exampleresourcegroup", "exampleserver", "XXXX-111",
            com.azure.core.util.Context.NONE);
    }
}
