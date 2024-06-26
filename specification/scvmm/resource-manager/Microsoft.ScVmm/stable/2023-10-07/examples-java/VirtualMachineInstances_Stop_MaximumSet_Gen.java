
import com.azure.resourcemanager.scvmm.models.SkipShutdown;
import com.azure.resourcemanager.scvmm.models.StopVirtualMachineOptions;

/**
 * Samples for VirtualMachineInstances Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/
     * VirtualMachineInstances_Stop_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineInstances_Stop_MaximumSet.
     * 
     * @param manager Entry point to ScvmmManager.
     */
    public static void virtualMachineInstancesStopMaximumSet(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualMachineInstances().stop("gtgclehcbsyave",
            new StopVirtualMachineOptions().withSkipShutdown(SkipShutdown.TRUE), com.azure.core.util.Context.NONE);
    }
}
