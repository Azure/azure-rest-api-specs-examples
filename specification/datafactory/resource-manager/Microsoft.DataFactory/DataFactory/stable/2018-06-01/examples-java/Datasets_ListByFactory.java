
/**
 * Samples for Datasets ListByFactory.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/Datasets_ListByFactory.json
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
