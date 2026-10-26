
import com.azure.resourcemanager.datatransfer.models.ResourceBody;

/**
 * Samples for Connections Link.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/linkConnection.json
     */
    /**
     * Sample code: Links the specified connection.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void linksTheSpecifiedConnection(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.connections().link("testRG", "testConnection", new ResourceBody().withId(
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.AzureDataTransfer/connections/testConnection"),
            com.azure.core.util.Context.NONE);
    }
}
