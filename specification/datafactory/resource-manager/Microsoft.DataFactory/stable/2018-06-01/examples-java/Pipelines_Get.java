
/**
 * Samples for Pipelines Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_Get.json
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
