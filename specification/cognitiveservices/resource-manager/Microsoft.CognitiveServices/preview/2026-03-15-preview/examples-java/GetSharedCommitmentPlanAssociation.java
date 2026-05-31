
/**
 * Samples for CommitmentPlans GetAssociation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/GetSharedCommitmentPlanAssociation.json
     */
    /**
     * Sample code: GetCommitmentPlan.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getCommitmentPlan(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.commitmentPlans().getAssociationWithResponse("resourceGroupName", "commitmentPlanName",
            "commitmentPlanAssociationName", com.azure.core.util.Context.NONE);
    }
}
