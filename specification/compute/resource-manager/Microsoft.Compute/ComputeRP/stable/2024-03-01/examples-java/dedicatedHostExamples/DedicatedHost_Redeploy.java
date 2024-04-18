
/**
 * Samples for DedicatedHosts Redeploy.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * dedicatedHostExamples/DedicatedHost_Redeploy.json
     */
    /**
     * Sample code: Redeploy Dedicated Host.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redeployDedicatedHost(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDedicatedHosts().redeploy("myResourceGroup",
            "myDedicatedHostGroup", "myHost", com.azure.core.util.Context.NONE);
    }
}
