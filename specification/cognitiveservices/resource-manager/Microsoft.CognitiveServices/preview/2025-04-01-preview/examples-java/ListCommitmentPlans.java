
/**
 * Samples for CommitmentPlans List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/
     * ListCommitmentPlans.json
     */
    /**
     * Sample code: ListCommitmentPlans.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listCommitmentPlans(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.commitmentPlans().list("resourceGroupName", "accountName", com.azure.core.util.Context.NONE);
    }
}
