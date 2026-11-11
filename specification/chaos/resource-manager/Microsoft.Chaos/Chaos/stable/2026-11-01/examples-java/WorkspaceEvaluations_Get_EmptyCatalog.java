
/**
 * Samples for WorkspaceEvaluations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/WorkspaceEvaluations_Get_EmptyCatalog.json
     */
    /**
     * Sample code: A valid empty published catalog succeeds with zero counts and no aggregate recommendation.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void aValidEmptyPublishedCatalogSucceedsWithZeroCountsAndNoAggregateRecommendation(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaceEvaluations().getWithResponse("exampleRG", "exampleWorkspace",
            "00000000-0000-0000-0000-000000000003", com.azure.core.util.Context.NONE);
    }
}
