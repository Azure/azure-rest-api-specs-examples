
/**
 * Samples for CommitmentPlans ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/ListSharedCommitmentPlansByResourceGroup.json
     */
    /**
     * Sample code: List Commitment Plans by Resource Group.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listCommitmentPlansByResourceGroup(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.commitmentPlans().listByResourceGroup("resourceGroupName", com.azure.core.util.Context.NONE);
    }
}
