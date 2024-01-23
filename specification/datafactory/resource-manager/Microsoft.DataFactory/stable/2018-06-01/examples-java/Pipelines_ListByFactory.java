
/**
 * Samples for Pipelines ListByFactory.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * Pipelines_ListByFactory.json
     */
    /**
     * Sample code: Pipelines_ListByFactory.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void pipelinesListByFactory(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.pipelines().listByFactory("exampleResourceGroup", "exampleFactoryName",
            com.azure.core.util.Context.NONE);
    }
}
