
import com.azure.resourcemanager.scvmm.models.DeleteFromHost;
import com.azure.resourcemanager.scvmm.models.ForceDelete;

/**
 * Samples for VirtualMachineInstances Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * VirtualMachineInstances_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineInstances_Delete_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void virtualMachineInstancesDeleteMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualMachineInstances().delete("gtgclehcbsyave", ForceDelete.TRUE, DeleteFromHost.TRUE,
            com.azure.core.util.Context.NONE);
    }
}
