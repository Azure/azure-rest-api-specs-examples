
/**
 * Samples for VirtualMachineInstances Start.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * VirtualMachineInstances_Start_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineInstances_Start_MinimumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void virtualMachineInstancesStartMinimumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualMachineInstances().start("gtgclehcbsyave", com.azure.core.util.Context.NONE);
    }
}
