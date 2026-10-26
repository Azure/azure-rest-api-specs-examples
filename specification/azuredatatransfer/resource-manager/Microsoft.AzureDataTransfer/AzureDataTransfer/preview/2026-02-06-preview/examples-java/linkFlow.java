
import com.azure.resourcemanager.datatransfer.models.ResourceBody;

/**
 * Samples for Flows Link.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/linkFlow.json
     */
    /**
     * Sample code: Links the specified flow.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void linksTheSpecifiedFlow(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.flows().link("testRG", "testConnection", "testFlow", new ResourceBody().withId(
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.AzureDataTransfer/connections/testConnection/flows/testFlow"),
            com.azure.core.util.Context.NONE);
    }
}
