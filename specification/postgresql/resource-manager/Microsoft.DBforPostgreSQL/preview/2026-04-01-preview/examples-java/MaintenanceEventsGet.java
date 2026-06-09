
/**
 * Samples for MaintenanceEvents Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/MaintenanceEventsGet.json
     */
    /**
     * Sample code: Get information about a maintenance event for a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void getInformationAboutAMaintenanceEventForAServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.maintenanceEvents().getWithResponse("exampleresourcegroup", "exampleserver", "XXXX-111",
            com.azure.core.util.Context.NONE);
    }
}
