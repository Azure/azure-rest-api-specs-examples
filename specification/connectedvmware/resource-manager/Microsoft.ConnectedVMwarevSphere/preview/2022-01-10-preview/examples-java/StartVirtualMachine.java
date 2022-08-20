import com.azure.core.util.Context;

/** Samples for VirtualMachines Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/StartVirtualMachine.json
     */
    /**
     * Sample code: StartVirtualMachine.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void startVirtualMachine(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.virtualMachines().start("testrg", "DemoVM", Context.NONE);
    }
}
