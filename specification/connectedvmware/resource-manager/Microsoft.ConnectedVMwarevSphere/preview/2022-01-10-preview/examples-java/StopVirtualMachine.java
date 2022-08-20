import com.azure.core.util.Context;
import com.azure.resourcemanager.connectedvmware.models.StopVirtualMachineOptions;

/** Samples for VirtualMachines Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/StopVirtualMachine.json
     */
    /**
     * Sample code: StopVirtualMachine.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void stopVirtualMachine(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager
            .virtualMachines()
            .stop("testrg", "DemoVM", new StopVirtualMachineOptions().withSkipShutdown(true), Context.NONE);
    }
}
