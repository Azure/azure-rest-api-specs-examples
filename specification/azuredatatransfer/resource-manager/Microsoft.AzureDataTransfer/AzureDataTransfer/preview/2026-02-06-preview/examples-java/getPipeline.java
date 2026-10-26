
/**
 * Samples for Pipelines GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/getPipeline.json
     */
    /**
     * Sample code: Gets pipeline resource.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void getsPipelineResource(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.pipelines().getByResourceGroupWithResponse("testRG", "testPipeline", com.azure.core.util.Context.NONE);
    }
}
