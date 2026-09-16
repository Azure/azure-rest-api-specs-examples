
/**
 * Samples for Entities List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/entities/GetEntities.json
     */
    /**
     * Sample code: Get all entities.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllEntities(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.entities().list("myRg", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}
