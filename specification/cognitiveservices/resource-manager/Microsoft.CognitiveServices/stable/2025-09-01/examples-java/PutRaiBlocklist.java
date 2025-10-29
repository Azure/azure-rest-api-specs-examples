
import com.azure.resourcemanager.cognitiveservices.models.RaiBlocklistProperties;

/**
 * Samples for RaiBlocklists CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
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
