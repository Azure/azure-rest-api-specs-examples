/** Samples for CommitmentTiers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-12-01/examples/ListCommitmentTiers.json
     */
    /**
     * Sample code: ListCommitmentTiers.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listCommitmentTiers(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.commitmentTiers().list("location", com.azure.core.util.Context.NONE);
    }
}
