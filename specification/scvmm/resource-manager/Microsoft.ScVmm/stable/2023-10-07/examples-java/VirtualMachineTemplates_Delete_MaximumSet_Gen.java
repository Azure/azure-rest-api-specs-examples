
import com.azure.resourcemanager.scvmm.models.ForceDelete;

/**
 * Samples for VirtualMachineTemplates Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * VirtualMachineTemplates_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineTemplates_Delete_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void virtualMachineTemplatesDeleteMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualMachineTemplates().delete("rgscvmm", "6", ForceDelete.TRUE, com.azure.core.util.Context.NONE);
    }
}
