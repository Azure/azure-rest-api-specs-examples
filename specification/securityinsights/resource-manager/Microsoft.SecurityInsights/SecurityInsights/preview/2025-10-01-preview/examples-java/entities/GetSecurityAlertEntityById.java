
/**
 * Samples for Entities Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/entities/GetSecurityAlertEntityById.json
     */
    /**
     * Sample code: Get a security alert entity.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getASecurityAlertEntity(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.entities().getWithResponse("myRg", "myWorkspace", "4aa486e0-6f85-41af-99ea-7acdce7be6c8",
            com.azure.core.util.Context.NONE);
    }
}
