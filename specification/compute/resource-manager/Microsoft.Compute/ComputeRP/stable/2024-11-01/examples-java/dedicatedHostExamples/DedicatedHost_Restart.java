
/**
 * Samples for DedicatedHosts Restart.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * dedicatedHostExamples/DedicatedHost_Restart.json
     */
    /**
     * Sample code: Restart Dedicated Host.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void restartDedicatedHost(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDedicatedHosts().restart("myResourceGroup",
            "myDedicatedHostGroup", "myHost", com.azure.core.util.Context.NONE);
    }
}
