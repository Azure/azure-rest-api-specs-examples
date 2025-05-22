
/**
 * Samples for CommitmentPlans GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/
     * GetSharedCommitmentPlan.json
     */
    /**
     * Sample code: Get Commitment Plan.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getCommitmentPlan(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.commitmentPlans().getByResourceGroupWithResponse("resourceGroupName", "commitmentPlanName",
            com.azure.core.util.Context.NONE);
    }
}
