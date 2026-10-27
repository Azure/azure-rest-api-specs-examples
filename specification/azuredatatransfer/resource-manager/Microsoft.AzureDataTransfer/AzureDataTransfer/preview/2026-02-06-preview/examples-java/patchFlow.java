
import com.azure.resourcemanager.datatransfer.models.FlowsPatch;

/**
 * Samples for Flows Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/patchFlow.json
     */
    /**
     * Sample code: Updates the flow resource.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void updatesTheFlowResource(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.flows().update("testRG", "testConnection", "testFlow", new FlowsPatch(),
            com.azure.core.util.Context.NONE);
    }
}
