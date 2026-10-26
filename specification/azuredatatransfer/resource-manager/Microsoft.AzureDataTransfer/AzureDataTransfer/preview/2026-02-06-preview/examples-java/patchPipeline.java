
import com.azure.resourcemanager.datatransfer.models.PipelinesPatch;

/**
 * Samples for Pipelines Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/patchPipeline.json
     */
    /**
     * Sample code: Updates the pipeline resource.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void updatesThePipelineResource(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.pipelines().update("testRG", "testPipeline", new PipelinesPatch(), com.azure.core.util.Context.NONE);
    }
}
