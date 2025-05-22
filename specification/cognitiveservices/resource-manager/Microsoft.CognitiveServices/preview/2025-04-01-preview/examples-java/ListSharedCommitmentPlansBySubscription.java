
/**
 * Samples for CommitmentPlans ListPlansBySubscription.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/
     * ListSharedCommitmentPlansBySubscription.json
     */
    /**
     * Sample code: List Accounts by Subscription.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listAccountsBySubscription(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.commitmentPlans().listPlansBySubscription(com.azure.core.util.Context.NONE);
    }
}
