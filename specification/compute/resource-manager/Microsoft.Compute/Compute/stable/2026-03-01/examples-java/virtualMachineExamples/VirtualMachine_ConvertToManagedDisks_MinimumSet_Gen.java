
/**
 * Samples for VirtualMachines ConvertToManagedDisks.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_ConvertToManagedDisks_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_ConvertToManagedDisks_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineConvertToManagedDisksMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().convertToManagedDisks("rgcompute", "aaaaaaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
