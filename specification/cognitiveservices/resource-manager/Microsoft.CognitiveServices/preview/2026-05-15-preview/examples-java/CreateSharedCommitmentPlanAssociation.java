
/**
 * Samples for CommitmentPlans CreateOrUpdateAssociation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/CreateSharedCommitmentPlanAssociation.json
     */
    /**
     * Sample code: PutCommitmentPlan.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void putCommitmentPlan(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.commitmentPlans().defineAssociation("commitmentPlanAssociationName")
            .withExistingCommitmentPlan("resourceGroupName", "commitmentPlanName")
            .withAccountId(
                "/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName")
            .create();
    }
}
