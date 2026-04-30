
/**
 * Samples for DedicatedHosts ListAvailableSizes.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/dedicatedHostExamples/DedicatedHost_ListAvailableSizes.json
     */
    /**
     * Sample code: Get Available Dedicated Host Sizes.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getAvailableDedicatedHostSizes(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDedicatedHosts().listAvailableSizes("myResourceGroup", "myDedicatedHostGroup",
            "myHost", com.azure.core.util.Context.NONE);
    }
}
