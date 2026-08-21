
/**
 * Samples for VirtualMachineScaleSetVMs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-04-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Get_WithCapacityReservation.json
     */
    /**
     * Sample code: Get VM scale set VM with capacity reservation.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getVMScaleSetVMWithCapacityReservation(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().getWithResponse("myResourceGroup", "{vmss-name}", "0",
            null, com.azure.core.util.Context.NONE);
    }
}
