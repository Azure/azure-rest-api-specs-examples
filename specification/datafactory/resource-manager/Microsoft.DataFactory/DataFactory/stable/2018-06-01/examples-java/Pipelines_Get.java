
/**
 * Samples for Pipelines Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/Pipelines_Get.json
     */
    /**
     * Sample code: Pipelines_Get.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void pipelinesGet(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.pipelines().getWithResponse("exampleResourceGroup", "exampleFactoryName", "examplePipeline", null,
            com.azure.core.util.Context.NONE);
    }
}
