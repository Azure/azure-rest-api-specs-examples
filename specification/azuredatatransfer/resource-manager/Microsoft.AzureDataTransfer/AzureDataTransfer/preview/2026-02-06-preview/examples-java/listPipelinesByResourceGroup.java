
/**
 * Samples for Pipelines ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/listPipelinesByResourceGroup.json
     */
    /**
     * Sample code: Gets pipelines in a resource group.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void
        getsPipelinesInAResourceGroup(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.pipelines().listByResourceGroup("testRG", com.azure.core.util.Context.NONE);
    }
}
