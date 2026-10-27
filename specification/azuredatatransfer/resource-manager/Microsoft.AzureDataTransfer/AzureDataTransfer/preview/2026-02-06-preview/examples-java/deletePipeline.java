
/**
 * Samples for Pipelines Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/deletePipeline.json
     */
    /**
     * Sample code: Deletes the pipeline resource.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void deletesThePipelineResource(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.pipelines().delete("testRG", "testPipeline", com.azure.core.util.Context.NONE);
    }
}
