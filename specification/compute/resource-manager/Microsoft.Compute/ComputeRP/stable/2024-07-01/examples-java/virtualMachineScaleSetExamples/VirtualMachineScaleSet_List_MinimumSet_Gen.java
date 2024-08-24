
/**
 * Samples for VirtualMachineScaleSets ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSet_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_List_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetListMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSets().listByResourceGroup("rgcompute",
            com.azure.core.util.Context.NONE);
    }
}
