import com.azure.core.util.Context;

/** Samples for DedicatedHosts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/dedicatedHostExamples/DedicatedHosts_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHosts_Delete_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dedicatedHostsDeleteMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDedicatedHosts()
            .delete("rgcompute", "aaaaaa", "aaaaaaaaaaaaaaa", Context.NONE);
    }
}
