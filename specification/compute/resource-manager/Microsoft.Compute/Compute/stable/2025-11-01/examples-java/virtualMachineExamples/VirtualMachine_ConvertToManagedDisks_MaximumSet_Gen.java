
/**
 * Samples for VirtualMachines ConvertToManagedDisks.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachine_ConvertToManagedDisks_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_ConvertToManagedDisks_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineConvertToManagedDisksMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().convertToManagedDisks("rgcompute", "aaaaaaa",
            com.azure.core.util.Context.NONE);
    }
}
