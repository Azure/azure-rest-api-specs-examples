
/**
 * Samples for PipelineRuns CancelSync.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/PipelineRuns_Cancel.
     * json
     */
    /**
     * Sample code: PipelineRuns_Cancel.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void pipelineRunsCancel(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.pipelineRuns().cancelWithResponse("exampleResourceGroup", "exampleFactoryName",
            "16ac5348-ff82-4f95-a80d-638c1d47b721", null, com.azure.core.util.Context.NONE);
    }
}
