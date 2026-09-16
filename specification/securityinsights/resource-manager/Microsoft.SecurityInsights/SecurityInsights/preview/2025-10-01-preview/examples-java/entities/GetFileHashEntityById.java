
/**
 * Samples for Entities Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/entities/GetFileHashEntityById.json
     */
    /**
     * Sample code: Get a file hash entity.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAFileHashEntity(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.entities().getWithResponse("myRg", "myWorkspace", "ea359fa6-c1e5-f878-e105-6344f3e399a1",
            com.azure.core.util.Context.NONE);
    }
}
