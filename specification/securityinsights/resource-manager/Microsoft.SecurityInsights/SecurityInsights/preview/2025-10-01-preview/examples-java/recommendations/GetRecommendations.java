
/**
 * Samples for GetRecommendations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/recommendations/GetRecommendations.json
     */
    /**
     * Sample code: Get Recommendations list.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        getRecommendationsList(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.getRecommendations().list("myRg", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}
