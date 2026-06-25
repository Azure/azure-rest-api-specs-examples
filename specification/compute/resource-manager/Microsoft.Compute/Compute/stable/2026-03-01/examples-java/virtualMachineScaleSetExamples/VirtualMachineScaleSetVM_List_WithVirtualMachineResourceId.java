
/**
 * Samples for VirtualMachineScaleSetVMs List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_List_WithVirtualMachineResourceId.json
     */
    /**
     * Sample code: List Vmss VMs with VirtualMachineResourceId.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        listVmssVMsWithVirtualMachineResourceId(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().list("resourceGroupname", "vmssName", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
