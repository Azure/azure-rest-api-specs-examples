
/**
 * Samples for CommitmentTiers List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/
     * ListCommitmentTiers.json
     */
    /**
     * Sample code: ListCommitmentTiers.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listCommitmentTiers(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.commitmentTiers().list("location", com.azure.core.util.Context.NONE);
    }
}
