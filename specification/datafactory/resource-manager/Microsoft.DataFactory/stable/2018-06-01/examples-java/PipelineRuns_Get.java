
/**
 * Samples for PipelineRuns GetSync.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/PipelineRuns_Get.json
     */
    /**
     * Sample code: PipelineRuns_Get.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void pipelineRunsGet(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.pipelineRuns().getWithResponse("exampleResourceGroup", "exampleFactoryName",
            "2f7fdb90-5df1-4b8e-ac2f-064cfa58202b", com.azure.core.util.Context.NONE);
    }
}
