
/**
 * Samples for PrivateEndPointConnections ListByFactory.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * PrivateEndPointConnections_ListByFactory.json
     */
    /**
     * Sample code: privateEndPointConnections_ListByFactory.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void
        privateEndPointConnectionsListByFactory(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.privateEndPointConnections().listByFactory("exampleResourceGroup", "exampleFactoryName",
            com.azure.core.util.Context.NONE);
    }
}
