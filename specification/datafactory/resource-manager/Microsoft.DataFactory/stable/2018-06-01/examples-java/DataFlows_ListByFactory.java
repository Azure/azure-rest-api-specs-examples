
/**
 * Samples for DataFlows ListByFactory.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * DataFlows_ListByFactory.json
     */
    /**
     * Sample code: DataFlows_ListByFactory.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void dataFlowsListByFactory(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.dataFlows().listByFactory("exampleResourceGroup", "exampleFactoryName",
            com.azure.core.util.Context.NONE);
    }
}
