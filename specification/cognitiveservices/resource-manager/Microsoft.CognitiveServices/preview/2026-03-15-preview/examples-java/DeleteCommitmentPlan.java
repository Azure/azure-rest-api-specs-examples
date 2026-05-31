
/**
 * Samples for CommitmentPlans Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/DeleteCommitmentPlan.json
     */
    /**
     * Sample code: DeleteCommitmentPlan.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        deleteCommitmentPlan(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.commitmentPlans().delete("resourceGroupName", "accountName", "commitmentPlanName",
            com.azure.core.util.Context.NONE);
    }
}
