
/**
 * Samples for CommitmentPlans DeleteAssociation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/DeleteSharedCommitmentPlanAssociation.json
     */
    /**
     * Sample code: DeleteCommitmentPlan.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        deleteCommitmentPlan(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.commitmentPlans().deleteAssociation("resourceGroupName", "commitmentPlanName",
            "commitmentPlanAssociationName", com.azure.core.util.Context.NONE);
    }
}
