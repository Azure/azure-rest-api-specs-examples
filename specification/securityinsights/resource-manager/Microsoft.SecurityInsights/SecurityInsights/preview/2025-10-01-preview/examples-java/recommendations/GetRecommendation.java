
/**
 * Samples for Get SingleRecommendation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/recommendations/GetRecommendation.json
     */
    /**
     * Sample code: Get a recommendation.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getARecommendation(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.gets().singleRecommendationWithResponse("myRg", "myWorkspace", "6d4b54eb-8684-4aa3-a156-3aa37b8014bc",
            com.azure.core.util.Context.NONE);
    }
}
