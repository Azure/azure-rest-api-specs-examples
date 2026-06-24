
/**
 * Samples for VirtualMachineScaleSets ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_ListBySubscription_ByLocation.json
     */
    /**
     * Sample code: Lists all the VM scale sets under the specified subscription for the specified location.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listsAllTheVMScaleSetsUnderTheSpecifiedSubscriptionForTheSpecifiedLocation(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().listByLocation("eastus", com.azure.core.util.Context.NONE);
    }
}
