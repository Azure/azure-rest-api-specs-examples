
/**
 * Samples for VirtualMachines List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_ListAll_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_ListAll_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineListAllMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().list(null, null, null, com.azure.core.util.Context.NONE);
    }
}
