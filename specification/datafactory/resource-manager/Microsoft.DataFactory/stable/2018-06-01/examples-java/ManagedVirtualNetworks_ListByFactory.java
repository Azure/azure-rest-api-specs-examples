
/**
 * Samples for ManagedVirtualNetworks ListByFactory.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * ManagedVirtualNetworks_ListByFactory.json
     */
    /**
     * Sample code: ManagedVirtualNetworks_ListByFactory.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void
        managedVirtualNetworksListByFactory(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.managedVirtualNetworks().listByFactory("exampleResourceGroup", "exampleFactoryName",
            com.azure.core.util.Context.NONE);
    }
}
