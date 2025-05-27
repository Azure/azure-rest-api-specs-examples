
/**
 * Samples for LinkedServices ListByFactory.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * LinkedServices_ListByFactory.json
     */
    /**
     * Sample code: LinkedServices_ListByFactory.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void linkedServicesListByFactory(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.linkedServices().listByFactory("exampleResourceGroup", "exampleFactoryName",
            com.azure.core.util.Context.NONE);
    }
}
