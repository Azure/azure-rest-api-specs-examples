
/**
 * Samples for VmInstanceHybridIdentityMetadatas ListByVirtualMachineInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * VmInstanceHybridIdentityMetadatas_ListByVirtualMachineInstance_MaximumSet_Gen.json
     */
    /**
     * Sample code: VmInstanceHybridIdentityMetadatas_ListByVirtualMachineInstance_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void vmInstanceHybridIdentityMetadatasListByVirtualMachineInstanceMaximumSet(
        com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.vmInstanceHybridIdentityMetadatas().listByVirtualMachineInstance("gtgclehcbsyave",
            com.azure.core.util.Context.NONE);
    }
}
