
/**
 * Samples for Pipelines Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Pipelines_Delete.json
     */
    /**
     * Sample code: Pipelines_Delete.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void pipelinesDelete(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.pipelines().deleteWithResponse("exampleResourceGroup", "exampleFactoryName", "examplePipeline",
            com.azure.core.util.Context.NONE);
    }
}
