
/**
 * Samples for VirtualMachineTemplates GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * VirtualMachineTemplates_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineTemplates_Get_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void virtualMachineTemplatesGetMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualMachineTemplates().getByResourceGroupWithResponse("rgscvmm", "4",
            com.azure.core.util.Context.NONE);
    }
}
