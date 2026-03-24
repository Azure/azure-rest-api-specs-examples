
/**
 * Samples for ManagedPrivateEndpoints ListByFactory.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/ManagedPrivateEndpoints_ListByFactory.json
     */
    /**
     * Sample code: ManagedPrivateEndpoints_ListByFactory.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void
        managedPrivateEndpointsListByFactory(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.managedPrivateEndpoints().listByFactory("exampleResourceGroup", "exampleFactoryName",
            "exampleManagedVirtualNetworkName", com.azure.core.util.Context.NONE);
    }
}
