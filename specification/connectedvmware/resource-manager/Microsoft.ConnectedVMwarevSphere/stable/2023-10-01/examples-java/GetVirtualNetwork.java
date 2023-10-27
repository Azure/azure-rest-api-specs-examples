/** Samples for VirtualNetworks GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/GetVirtualNetwork.json
     */
    /**
     * Sample code: GetVirtualNetwork.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void getVirtualNetwork(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager
            .virtualNetworks()
            .getByResourceGroupWithResponse("testrg", "ProdNetwork", com.azure.core.util.Context.NONE);
    }
}
