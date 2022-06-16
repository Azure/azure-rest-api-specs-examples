import com.azure.core.util.Context;

/** Samples for PipelineJobs List. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-job-list.json
     */
    /**
     * Sample code: List all pipeline jobs.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void listAllPipelineJobs(com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager.pipelineJobs().list("testrg", "testaccount2", null, 2, Context.NONE);
    }
}
