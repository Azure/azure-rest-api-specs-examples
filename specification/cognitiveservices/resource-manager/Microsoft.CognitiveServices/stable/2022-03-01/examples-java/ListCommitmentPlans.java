import com.azure.core.util.Context;

/** Samples for CommitmentPlans List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/ListCommitmentPlans.json
     */
    /**
     * Sample code: ListCommitmentPlans.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listCommitmentPlans(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.commitmentPlans().list("resourceGroupName", "accountName", Context.NONE);
    }
}
