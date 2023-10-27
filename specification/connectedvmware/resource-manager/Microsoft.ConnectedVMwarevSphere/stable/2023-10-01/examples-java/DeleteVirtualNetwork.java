/** Samples for VirtualNetworks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/DeleteVirtualNetwork.json
     */
    /**
     * Sample code: DeleteVirtualNetwork.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void deleteVirtualNetwork(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.virtualNetworks().delete("testrg", "ProdNetwork", null, com.azure.core.util.Context.NONE);
    }
}
