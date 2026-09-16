
/**
 * Samples for Entities Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/entities/GetFileEntityById.json
     */
    /**
     * Sample code: Get a file entity.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAFileEntity(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.entities().getWithResponse("myRg", "myWorkspace", "af378b21-b4aa-4fe7-bc70-13f8621a322f",
            com.azure.core.util.Context.NONE);
    }
}
