
/**
 * Samples for VirtualMachineScaleSetVMs GetInstanceView.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-04-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Get_InstanceViewWithCapacityReservationType.
     * json
     */
    /**
     * Sample code: Get instance view of a virtual machine from a VM scale set that is eligible for and consuming an
     * open capacity reservation.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getInstanceViewOfAVirtualMachineFromAVMScaleSetThatIsEligibleForAndConsumingAnOpenCapacityReservation(
            com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().getInstanceViewWithResponse("myResourceGroup",
            "myVirtualMachineScaleSet", "0", com.azure.core.util.Context.NONE);
    }
}
