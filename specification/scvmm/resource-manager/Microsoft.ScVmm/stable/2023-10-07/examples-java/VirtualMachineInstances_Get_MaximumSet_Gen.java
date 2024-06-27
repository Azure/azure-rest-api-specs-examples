
/**
 * Samples for VirtualMachineInstances Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * VirtualMachineInstances_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineInstances_Get_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void virtualMachineInstancesGetMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualMachineInstances().getWithResponse("gtgclehcbsyave", com.azure.core.util.Context.NONE);
    }
}
