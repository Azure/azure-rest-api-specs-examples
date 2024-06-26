
/**
 * Samples for VirtualMachineInstances Restart.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * VirtualMachineInstances_Restart_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineInstances_Restart_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void virtualMachineInstancesRestartMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualMachineInstances().restart("gtgclehcbsyave", com.azure.core.util.Context.NONE);
    }
}
