import com.azure.core.util.Context;

/** Samples for VirtualMachineTemplates List. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/ListVirtualMachineTemplates.json
     */
    /**
     * Sample code: ListVirtualMachineTemplates.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listVirtualMachineTemplates(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.virtualMachineTemplates().list(Context.NONE);
    }
}
