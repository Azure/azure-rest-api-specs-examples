
/**
 * Samples for VirtualMachineInstances Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * VirtualMachineInstances_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineInstances_Delete_MinimumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void virtualMachineInstancesDeleteMinimumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualMachineInstances().delete("gtgclehcbsyave", null, null, com.azure.core.util.Context.NONE);
    }
}
