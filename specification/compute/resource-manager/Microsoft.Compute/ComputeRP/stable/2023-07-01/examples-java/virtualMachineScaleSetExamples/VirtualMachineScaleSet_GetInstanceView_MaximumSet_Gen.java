/** Samples for VirtualMachineScaleSets GetInstanceView. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_GetInstanceView_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_GetInstanceView_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetGetInstanceViewMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .getInstanceViewWithResponse("rgcompute", "aaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE);
    }
}
