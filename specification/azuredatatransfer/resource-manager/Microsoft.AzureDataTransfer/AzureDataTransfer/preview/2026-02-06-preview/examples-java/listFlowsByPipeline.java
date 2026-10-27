
import com.azure.resourcemanager.datatransfer.models.ConnectionIdList;
import java.util.Arrays;

/**
 * Samples for Pipelines List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/listFlowsByPipeline.json
     */
    /**
     * Sample code: Lists flows for a pipeline.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void listsFlowsForAPipeline(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.pipelines().listWithResponse("testRG", "testPipeline", new ConnectionIdList().withValue(Arrays.asList(
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.AzureDataTransfer/connections/testConnection")),
            com.azure.core.util.Context.NONE);
    }
}
