
import com.azure.resourcemanager.datatransfer.models.Action;
import com.azure.resourcemanager.datatransfer.models.DateTransferActionType;
import com.azure.resourcemanager.datatransfer.models.TargetType;
import java.util.Arrays;

/**
 * Samples for Pipelines ExecuteAction.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/executeAction.json
     */
    /**
     * Sample code: Executes action for a pipeline.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void executesActionForAPipeline(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.pipelines().executeAction("testRG", "testPipeline", new Action()
            .withActionType(DateTransferActionType.FORCE_DISABLE).withTargetType(TargetType.CONNECTION)
            .withTargets(Arrays.asList(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.AzureDataTransfer/connections/testConnection")),
            com.azure.core.util.Context.NONE);
    }
}
