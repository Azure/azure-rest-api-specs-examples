
/**
 * Samples for Entities Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/entities/GetProcessEntityById.json
     */
    /**
     * Sample code: Get a process entity.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAProcessEntity(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.entities().getWithResponse("myRg", "myWorkspace", "7264685c-038c-42c6-948c-38e14ef1fb98",
            com.azure.core.util.Context.NONE);
    }
}
