import com.azure.core.util.Context;

/** Samples for VirtualMachines Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/DeleteVirtualMachine.json
     */
    /**
     * Sample code: DeleteVirtualMachine.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void deleteVirtualMachine(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.virtualMachines().delete("testrg", "DemoVM", null, null, Context.NONE);
    }
}
