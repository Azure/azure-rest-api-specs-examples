
/**
 * Samples for CommitmentPlans ListAssociations.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/ListSharedCommitmentPlanAssociations.json
     */
    /**
     * Sample code: ListCommitmentPlans.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listCommitmentPlans(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.commitmentPlans().listAssociations("resourceGroupName", "commitmentPlanName",
            com.azure.core.util.Context.NONE);
    }
}
