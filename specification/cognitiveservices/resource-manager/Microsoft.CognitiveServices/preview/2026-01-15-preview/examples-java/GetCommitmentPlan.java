
/**
 * Samples for CommitmentPlans Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/GetCommitmentPlan.json
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
