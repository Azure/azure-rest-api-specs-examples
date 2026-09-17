
/**
 * Samples for Hunts List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/hunts/GetHunts.json
     */
    /**
     * Sample code: Get all hunts.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllHunts(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.hunts().list("myRg", "myWorkspace", null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
