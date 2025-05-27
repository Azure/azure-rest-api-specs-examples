
/**
 * Samples for Datasets ListByFactory.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * Datasets_ListByFactory.json
     */
    /**
     * Sample code: Datasets_ListByFactory.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void datasetsListByFactory(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.datasets().listByFactory("exampleResourceGroup", "exampleFactoryName",
            com.azure.core.util.Context.NONE);
    }
}
