
/**
 * Samples for HuntComments List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/hunts/GetHuntComments.json
     */
    /**
     * Sample code: Get all hunt comments.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllHuntComments(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.huntComments().list("myRg", "myWorkspace", "163e7b2a-a2ec-4041-aaba-d878a38f265f", null, null, null,
            null, com.azure.core.util.Context.NONE);
    }
}
