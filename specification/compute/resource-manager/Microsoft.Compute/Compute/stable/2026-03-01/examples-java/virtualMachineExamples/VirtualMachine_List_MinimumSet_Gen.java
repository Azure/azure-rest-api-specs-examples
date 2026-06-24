
/**
 * Samples for VirtualMachines ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_List_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineListMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().listByResourceGroup("rgcompute", null, null,
            com.azure.core.util.Context.NONE);
    }
}
