
import com.azure.resourcemanager.cognitiveservices.models.RaiBlocklistProperties;

/**
 * Samples for RaiBlocklists CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/
     * PutRaiBlocklist.json
     */
    /**
     * Sample code: PutRaiBlocklist.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void putRaiBlocklist(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiBlocklists().define("raiBlocklistName").withExistingAccount("resourceGroupName", "accountName")
            .withProperties(new RaiBlocklistProperties().withDescription("Basic blocklist description")).create();
    }
}
