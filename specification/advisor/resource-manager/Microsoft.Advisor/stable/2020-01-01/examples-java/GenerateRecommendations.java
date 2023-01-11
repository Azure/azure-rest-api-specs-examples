/** Samples for Recommendations Generate. */
public final class Main {
    /*
     * x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/GenerateRecommendations.json
     */
    /**
     * Sample code: GenerateRecommendations.
     *
     * @param manager Entry point to AdvisorManager.
     */
    public static void generateRecommendations(com.azure.resourcemanager.advisor.AdvisorManager manager) {
        manager.recommendations().generateWithResponse(com.azure.core.util.Context.NONE);
    }
}
