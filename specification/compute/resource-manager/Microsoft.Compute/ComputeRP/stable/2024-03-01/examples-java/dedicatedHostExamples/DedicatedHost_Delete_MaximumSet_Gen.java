
/**
 * Samples for DedicatedHosts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * dedicatedHostExamples/DedicatedHost_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHost_Delete_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dedicatedHostDeleteMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDedicatedHosts().delete("rgcompute", "aaaaaa",
            "aaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
