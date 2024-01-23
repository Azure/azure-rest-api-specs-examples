
/**
 * Samples for ManagedPrivateEndpoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * ManagedPrivateEndpoints_Delete.json
     */
    /**
     * Sample code: ManagedVirtualNetworks_Delete.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void managedVirtualNetworksDelete(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.managedPrivateEndpoints().deleteWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleManagedVirtualNetworkName", "exampleManagedPrivateEndpointName", com.azure.core.util.Context.NONE);
    }
}
