
/**
 * Samples for DedicatedHosts Redeploy.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/dedicatedHostExamples/DedicatedHost_Redeploy.json
     */
    /**
     * Sample code: Redeploy Dedicated Host.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void redeployDedicatedHost(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDedicatedHosts().redeploy("myResourceGroup", "myDedicatedHostGroup", "myHost",
            com.azure.core.util.Context.NONE);
    }
}
