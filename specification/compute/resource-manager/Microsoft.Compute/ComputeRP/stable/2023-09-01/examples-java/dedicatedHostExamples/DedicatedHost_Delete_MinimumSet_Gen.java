
/**
 * Samples for DedicatedHosts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/
     * dedicatedHostExamples/DedicatedHost_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHost_Delete_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dedicatedHostDeleteMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDedicatedHosts().delete("rgcompute", "aaaaaaaaaaaaaaa",
            "aaaaa", com.azure.core.util.Context.NONE);
    }
}
