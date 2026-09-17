
import com.azure.resourcemanager.securityinsights.models.RecommendationPatch;
import com.azure.resourcemanager.securityinsights.models.RecommendationPatchProperties;
import com.azure.resourcemanager.securityinsights.models.State;

/**
 * Samples for Update Recommendation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/recommendations/PatchRecommendation.json
     */
    /**
     * Sample code: Creates a recommendation.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        createsARecommendation(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.updates().recommendationWithResponse("myRg", "myWorkspace", "6d4b54eb-8684-4aa3-a156-3aa37b8014bc",
            new RecommendationPatch().withProperties(new RecommendationPatchProperties().withState(State.ACTIVE)),
            com.azure.core.util.Context.NONE);
    }
}
