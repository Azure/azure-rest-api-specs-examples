
import com.azure.resourcemanager.cognitiveservices.models.RaiBlocklistItemBulkRequest;
import com.azure.resourcemanager.cognitiveservices.models.RaiBlocklistItemProperties;
import java.util.Arrays;

/**
 * Samples for RaiBlocklistItems BatchAdd.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/
     * AddRaiBlocklistItems.json
     */
    /**
     * Sample code: AddRaiBlocklistItems.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        addRaiBlocklistItems(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiBlocklistItems().batchAddWithResponse("resourceGroupName", "accountName", "myblocklist",
            Arrays.asList(
                new RaiBlocklistItemBulkRequest().withName("myblocklistitem1").withProperties(
                    new RaiBlocklistItemProperties().withPattern("^[a-z0-9_-]{2,16}$").withIsRegex(true)),
                new RaiBlocklistItemBulkRequest().withName("myblocklistitem2")
                    .withProperties(new RaiBlocklistItemProperties().withPattern("blockwords").withIsRegex(false))),
            com.azure.core.util.Context.NONE);
    }
}
