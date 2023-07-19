/** Samples for CommitmentPlans DeleteAssociation. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/DeleteSharedCommitmentPlanAssociation.json
     */
    /**
     * Sample code: DeleteCommitmentPlan.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void deleteCommitmentPlan(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager
            .commitmentPlans()
            .deleteAssociation(
                "resourceGroupName",
                "commitmentPlanName",
                "commitmentPlanAssociationName",
                com.azure.core.util.Context.NONE);
    }
}
