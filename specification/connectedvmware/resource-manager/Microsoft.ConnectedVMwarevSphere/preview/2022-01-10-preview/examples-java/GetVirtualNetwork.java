import com.azure.core.util.Context;

/** Samples for VirtualNetworks GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/GetVirtualNetwork.json
     */
    /**
     * Sample code: GetVirtualNetwork.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void getVirtualNetwork(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.virtualNetworks().getByResourceGroupWithResponse("testrg", "ProdNetwork", Context.NONE);
    }
}
