
/**
 * Samples for VirtualMachineScaleSets Redeploy.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSet_Redeploy_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_Redeploy_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineScaleSetRedeployMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSets().redeploy("rgcompute",
            "aaaaaaaaaaaaaaaaaaaaaa", null, com.azure.core.util.Context.NONE);
    }
}
