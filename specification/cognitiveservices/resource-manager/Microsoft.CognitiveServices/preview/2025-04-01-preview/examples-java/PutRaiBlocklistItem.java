
import com.azure.resourcemanager.cognitiveservices.models.RaiBlocklistItemProperties;

/**
 * Samples for RaiBlocklistItems CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/
     * PutRaiBlocklistItem.json
     */
    /**
     * Sample code: PutRaiBlocklistItem.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        putRaiBlocklistItem(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiBlocklistItems().define("raiBlocklistItemName")
            .withExistingRaiBlocklist("resourceGroupName", "accountName", "raiBlocklistName")
            .withProperties(new RaiBlocklistItemProperties().withPattern("Pattern To Block").withIsRegex(false))
            .create();
    }
}
