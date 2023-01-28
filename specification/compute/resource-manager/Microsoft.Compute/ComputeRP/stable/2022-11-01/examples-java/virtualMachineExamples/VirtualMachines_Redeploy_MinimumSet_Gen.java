/** Samples for VirtualMachines Redeploy. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/virtualMachineExamples/VirtualMachines_Redeploy_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachines_Redeploy_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachinesRedeployMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .redeploy("rgcompute", "aaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
