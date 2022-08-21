import com.azure.core.util.Context;

/** Samples for VirtualMachines Restart. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/RestartVirtualMachine.json
     */
    /**
     * Sample code: RestartVirtualMachine.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void restartVirtualMachine(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.virtualMachines().restart("testrg", "DemoVM", Context.NONE);
    }
}
