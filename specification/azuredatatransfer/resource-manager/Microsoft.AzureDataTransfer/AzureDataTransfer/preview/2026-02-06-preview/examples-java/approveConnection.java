
import com.azure.resourcemanager.datatransfer.models.ResourceBody;

/**
 * Samples for Pipelines ApproveConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/approveConnection.json
     */
    /**
     * Sample code: Approves the specified connection in a pipeline.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void
        approvesTheSpecifiedConnectionInAPipeline(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.pipelines().approveConnection("testRG", "testPipeline", new ResourceBody().withId(
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.AzureDataTransfer/connections/testConnection")
            .withStatusReason("Example reason"), com.azure.core.util.Context.NONE);
    }
}
