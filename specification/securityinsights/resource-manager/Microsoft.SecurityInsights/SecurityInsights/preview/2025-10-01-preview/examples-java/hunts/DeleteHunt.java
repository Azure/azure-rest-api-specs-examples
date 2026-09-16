
/**
 * Samples for Hunts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/hunts/DeleteHunt.json
     */
    /**
     * Sample code: Delete a hunt.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteAHunt(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.hunts().deleteWithResponse("myRg", "myWorkspace", "163e7b2a-a2ec-4041-aaba-d878a38f265f",
            com.azure.core.util.Context.NONE);
    }
}
