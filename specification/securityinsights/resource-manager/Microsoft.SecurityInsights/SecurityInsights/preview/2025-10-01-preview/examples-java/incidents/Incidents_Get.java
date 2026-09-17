
/**
 * Samples for Incidents Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/incidents/Incidents_Get.json
     */
    /**
     * Sample code: Incidents_Get.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void incidentsGet(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.incidents().getWithResponse("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
            com.azure.core.util.Context.NONE);
    }
}
