
/**
 * Samples for VirtualMachines PowerOff.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_PowerOff_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_PowerOff_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachinePowerOffMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().powerOff("rgcompute", "aaaaaaaaaaaaaaaaaa", null,
            com.azure.core.util.Context.NONE);
    }
}
