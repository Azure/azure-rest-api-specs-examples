/** Samples for RecommendationMetadata Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/GetRecommendationMetadataEntity.json
     */
    /**
     * Sample code: GetMetadata.
     *
     * @param manager Entry point to AdvisorManager.
     */
    public static void getMetadata(com.azure.resourcemanager.advisor.AdvisorManager manager) {
        manager.recommendationMetadatas().getWithResponse("types", com.azure.core.util.Context.NONE);
    }
}
