
/**
 * Samples for VirtualMachines Redeploy.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * virtualMachineExamples/VirtualMachine_Redeploy_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_Redeploy_MaximumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineRedeployMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachines().redeploy("rgcompute", "a",
            com.azure.core.util.Context.NONE);
    }
}
