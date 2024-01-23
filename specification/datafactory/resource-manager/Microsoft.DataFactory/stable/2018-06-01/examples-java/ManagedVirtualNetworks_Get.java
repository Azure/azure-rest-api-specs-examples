
/**
 * Samples for ManagedVirtualNetworks Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * ManagedVirtualNetworks_Get.json
     */
    /**
     * Sample code: ManagedVirtualNetworks_Get.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void managedVirtualNetworksGet(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.managedVirtualNetworks().getWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleManagedVirtualNetworkName", null, com.azure.core.util.Context.NONE);
    }
}
