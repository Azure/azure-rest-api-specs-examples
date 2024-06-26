
/**
 * Samples for VirtualMachineInstances List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * VirtualMachineInstances_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineInstances_List_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void virtualMachineInstancesListMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualMachineInstances().list("gtgclehcbsyave", com.azure.core.util.Context.NONE);
    }
}
