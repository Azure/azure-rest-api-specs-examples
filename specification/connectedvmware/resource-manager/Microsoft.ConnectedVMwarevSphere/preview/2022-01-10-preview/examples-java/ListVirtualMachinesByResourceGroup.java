import com.azure.core.util.Context;

/** Samples for VirtualMachines ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/ListVirtualMachinesByResourceGroup.json
     */
    /**
     * Sample code: ListVirtualMachinesByResourceGroup.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listVirtualMachinesByResourceGroup(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.virtualMachines().listByResourceGroup("testrg", Context.NONE);
    }
}
