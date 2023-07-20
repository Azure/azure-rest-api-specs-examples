/** Samples for CommitmentPlans GetAssociation. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/GetSharedCommitmentPlanAssociation.json
     */
    /**
     * Sample code: GetCommitmentPlan.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getCommitmentPlan(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager
            .commitmentPlans()
            .getAssociationWithResponse(
                "resourceGroupName",
                "commitmentPlanName",
                "commitmentPlanAssociationName",
                com.azure.core.util.Context.NONE);
    }
}
