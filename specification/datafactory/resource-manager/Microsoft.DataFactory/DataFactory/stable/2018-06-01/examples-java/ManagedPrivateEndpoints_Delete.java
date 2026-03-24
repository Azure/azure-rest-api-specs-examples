
/**
 * Samples for ManagedPrivateEndpoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/ManagedPrivateEndpoints_Delete.json
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
