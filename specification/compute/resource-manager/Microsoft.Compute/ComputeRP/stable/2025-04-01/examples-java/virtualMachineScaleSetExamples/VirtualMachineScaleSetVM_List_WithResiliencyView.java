
/**
 * Samples for VirtualMachineScaleSetVMs List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_List_WithResiliencyView.json
     */
    /**
     * Sample code: List Vmss VMs with ResilientVMDeletionStatus.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listVmssVMsWithResilientVMDeletionStatus(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetVMs().list("resourceGroupname",
            "vmssName", null, null, null, com.azure.core.util.Context.NONE);
    }
}
