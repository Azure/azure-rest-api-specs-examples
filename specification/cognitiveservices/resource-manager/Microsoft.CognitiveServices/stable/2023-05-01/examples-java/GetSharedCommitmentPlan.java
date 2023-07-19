/** Samples for CommitmentPlans GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/GetSharedCommitmentPlan.json
     */
    /**
     * Sample code: Get Commitment Plan.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getCommitmentPlan(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager
            .commitmentPlans()
            .getByResourceGroupWithResponse(
                "resourceGroupName", "commitmentPlanName", com.azure.core.util.Context.NONE);
    }
}
