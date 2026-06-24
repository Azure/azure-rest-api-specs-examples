
/**
 * Samples for VirtualMachines List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_ListAll_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_ListAll_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineListAllMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().list("aaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaa", null,
            com.azure.core.util.Context.NONE);
    }
}
