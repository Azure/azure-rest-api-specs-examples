
/**
 * Samples for ManagedUnsupportedVMSizes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/managedUnsupportedVMSizesList_example.json
     */
    /**
     * Sample code: List unsupported vm sizes.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void listUnsupportedVmSizes(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.managedUnsupportedVMSizes().list("eastus", com.azure.core.util.Context.NONE);
    }
}
