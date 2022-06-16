import com.azure.core.util.Context;
import com.azure.resourcemanager.scvmm.models.StopVirtualMachineOptions;

/** Samples for VirtualMachines Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/StopVirtualMachine.json
     */
    /**
     * Sample code: StopVirtualMachine.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void stopVirtualMachine(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager
            .virtualMachines()
            .stop("testrg", "DemoVM", new StopVirtualMachineOptions().withSkipShutdown(true), Context.NONE);
    }
}
