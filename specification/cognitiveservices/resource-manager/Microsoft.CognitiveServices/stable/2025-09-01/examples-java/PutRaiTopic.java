
import com.azure.resourcemanager.cognitiveservices.models.RaiTopicProperties;

/**
 * Samples for RaiTopics CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * PutRaiTopic.json
     */
    /**
     * Sample code: PutRaiTopic.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void putRaiTopic(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiTopics().define("raiTopicName").withExistingAccount("resourceGroupName", "accountName")
            .withProperties(
                new RaiTopicProperties().withTopicName("raiTopicName").withDescription("This is a sample topic.")
                    .withSampleBlobUrl("https://example.blob.core.windows.net/sampleblob"))
            .create();
    }
}
