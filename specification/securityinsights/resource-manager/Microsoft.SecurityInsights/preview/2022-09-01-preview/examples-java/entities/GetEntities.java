
/**
 * Samples for Entities List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * entities/GetEntities.json
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
