
/**
 * Samples for VirtualMachineTemplates Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * VirtualMachineTemplates_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineTemplates_Delete_MinimumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void virtualMachineTemplatesDeleteMinimumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualMachineTemplates().delete("rgscvmm", "5", null, com.azure.core.util.Context.NONE);
    }
}
