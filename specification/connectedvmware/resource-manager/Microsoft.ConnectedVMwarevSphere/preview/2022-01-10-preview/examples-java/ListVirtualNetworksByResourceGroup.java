import com.azure.core.util.Context;

/** Samples for VirtualNetworks ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/ListVirtualNetworksByResourceGroup.json
     */
    /**
     * Sample code: ListVirtualNetworksByResourceGroup.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listVirtualNetworksByResourceGroup(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.virtualNetworks().listByResourceGroup("testrg", Context.NONE);
    }
}
