import com.azure.core.util.Context;

/** Samples for CommitmentPlans Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-10-01/examples/DeleteCommitmentPlan.json
     */
    /**
     * Sample code: DeleteCommitmentPlan.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void deleteCommitmentPlan(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.commitmentPlans().delete("resourceGroupName", "accountName", "commitmentPlanName", Context.NONE);
    }
}
