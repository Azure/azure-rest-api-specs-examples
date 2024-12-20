
/**
 * Samples for ManagedUnsupportedVMSizes Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-06-01-preview/
     * examples/managedUnsupportedVMSizesGet_example.json
     */
    /**
     * Sample code: Get unsupported vm sizes.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void getUnsupportedVmSizes(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.managedUnsupportedVMSizes().getWithResponse("eastus", "Standard_B1ls1",
            com.azure.core.util.Context.NONE);
    }
}
