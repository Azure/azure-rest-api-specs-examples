/** Samples for Recommendations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/ListRecommendations.json
     */
    /**
     * Sample code: ListRecommendations.
     *
     * @param manager Entry point to AdvisorManager.
     */
    public static void listRecommendations(com.azure.resourcemanager.advisor.AdvisorManager manager) {
        manager.recommendations().list(null, 10, null, com.azure.core.util.Context.NONE);
    }
}
