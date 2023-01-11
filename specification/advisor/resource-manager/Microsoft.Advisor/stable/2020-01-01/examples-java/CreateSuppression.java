/** Samples for Suppressions Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/CreateSuppression.json
     */
    /**
     * Sample code: CreateSuppression.
     *
     * @param manager Entry point to AdvisorManager.
     */
    public static void createSuppression(com.azure.resourcemanager.advisor.AdvisorManager manager) {
        manager
            .suppressions()
            .define("suppressionName1")
            .withExistingRecommendation("resourceUri", "recommendationId")
            .withTtl("07:00:00:00")
            .create();
    }
}
