
/**
 * Samples for GlobalParameters ListByFactory.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/GlobalParameters_ListByFactory.json
     */
    /**
     * Sample code: GlobalParameters_ListByFactory.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void globalParametersListByFactory(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.globalParameters().listByFactory("exampleResourceGroup", "exampleFactoryName",
            com.azure.core.util.Context.NONE);
    }
}
