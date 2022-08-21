import com.azure.core.util.Context;

/** Samples for VirtualMachines List. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/ListVirtualMachines.json
     */
    /**
     * Sample code: ListVirtualMachines.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listVirtualMachines(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.virtualMachines().list(Context.NONE);
    }
}
