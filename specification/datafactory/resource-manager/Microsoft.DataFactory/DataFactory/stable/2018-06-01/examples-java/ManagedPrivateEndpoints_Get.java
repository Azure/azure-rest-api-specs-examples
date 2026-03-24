
/**
 * Samples for ManagedPrivateEndpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/ManagedPrivateEndpoints_Get.json
     */
    /**
     * Sample code: ManagedPrivateEndpoints_Get.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void managedPrivateEndpointsGet(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.managedPrivateEndpoints().getWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleManagedVirtualNetworkName", "exampleManagedPrivateEndpointName", null,
            com.azure.core.util.Context.NONE);
    }
}
