
/**
 * Samples for DedicatedHosts Restart.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/dedicatedHostExamples/DedicatedHost_Restart.json
     */
    /**
     * Sample code: Restart Dedicated Host.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void restartDedicatedHost(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDedicatedHosts().restart("myResourceGroup", "myDedicatedHostGroup", "myHost",
            com.azure.core.util.Context.NONE);
    }
}
