
import com.azure.resourcemanager.scvmm.models.VirtualMachineInstanceUpdate;

/**
 * Samples for VirtualMachineInstances Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * VirtualMachineInstances_Update_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineInstances_Update_MinimumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void virtualMachineInstancesUpdateMinimumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualMachineInstances().update("gtgclehcbsyave", new VirtualMachineInstanceUpdate(),
            com.azure.core.util.Context.NONE);
    }
}
