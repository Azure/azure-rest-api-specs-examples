
/**
 * Samples for VirtualMachineScaleSetLifeCycleHookEvents List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetLifeCycleHookEvent_List.json
     */
    /**
     * Sample code: Gets a list of all lifecycle hook events in a virtual machine scale set.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getsAListOfAllLifecycleHookEventsInAVirtualMachineScaleSet(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetLifeCycleHookEvents().list("RG01", "VMSS01",
            com.azure.core.util.Context.NONE);
    }
}
