
/**
 * Samples for CommitmentPlans Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/
     * GetCommitmentPlan.json
     */
    /**
     * Sample code: GetCommitmentPlan.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getCommitmentPlan(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.commitmentPlans().getWithResponse("resourceGroupName", "accountName", "commitmentPlanName",
            com.azure.core.util.Context.NONE);
    }
}
