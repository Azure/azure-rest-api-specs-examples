
/**
 * Samples for CommitmentPlans Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/
     * DeleteCommitmentPlan.json
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
