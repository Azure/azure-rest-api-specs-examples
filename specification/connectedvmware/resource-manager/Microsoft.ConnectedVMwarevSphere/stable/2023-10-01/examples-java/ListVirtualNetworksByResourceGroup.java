/** Samples for VirtualNetworks ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/ListVirtualNetworksByResourceGroup.json
     */
    /**
     * Sample code: ListVirtualNetworksByResourceGroup.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listVirtualNetworksByResourceGroup(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.virtualNetworks().listByResourceGroup("testrg", com.azure.core.util.Context.NONE);
    }
}
