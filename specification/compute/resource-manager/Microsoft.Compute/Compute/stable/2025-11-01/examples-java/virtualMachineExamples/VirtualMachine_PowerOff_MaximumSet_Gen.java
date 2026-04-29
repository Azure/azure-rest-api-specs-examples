
/**
 * Samples for VirtualMachines PowerOff.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachine_PowerOff_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_PowerOff_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachinePowerOffMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().powerOff("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaaaa", true,
            com.azure.core.util.Context.NONE);
    }
}
