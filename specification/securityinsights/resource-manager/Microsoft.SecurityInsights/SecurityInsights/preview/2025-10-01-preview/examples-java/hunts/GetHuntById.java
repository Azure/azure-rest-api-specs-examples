
/**
 * Samples for Hunts Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/hunts/GetHuntById.json
     */
    /**
     * Sample code: Get a hunt.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAHunt(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.hunts().getWithResponse("myRg", "myWorkspace", "163e7b2a-a2ec-4041-aaba-d878a38f265f",
            com.azure.core.util.Context.NONE);
    }
}
