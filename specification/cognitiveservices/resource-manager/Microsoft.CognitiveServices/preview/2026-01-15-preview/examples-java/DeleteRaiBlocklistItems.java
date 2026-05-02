
import java.util.Arrays;

/**
 * Samples for RaiBlocklistItems BatchDelete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/DeleteRaiBlocklistItems.json
     */
    /**
     * Sample code: DeleteRaiBlocklistItems.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        deleteRaiBlocklistItems(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiBlocklistItems().batchDeleteWithResponse("resourceGroupName", "accountName", "raiBlocklistName",
            Arrays.asList("myblocklistitem1", "myblocklistitem2"), com.azure.core.util.Context.NONE);
    }
}
