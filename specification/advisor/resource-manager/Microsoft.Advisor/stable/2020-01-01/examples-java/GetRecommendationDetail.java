/** Samples for Recommendations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/GetRecommendationDetail.json
     */
    /**
     * Sample code: GetRecommendationDetail.
     *
     * @param manager Entry point to AdvisorManager.
     */
    public static void getRecommendationDetail(com.azure.resourcemanager.advisor.AdvisorManager manager) {
        manager.recommendations().getWithResponse("resourceUri", "recommendationId", com.azure.core.util.Context.NONE);
    }
}
