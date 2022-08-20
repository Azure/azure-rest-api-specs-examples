import com.azure.core.util.Context;

/** Samples for VirtualNetworks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/DeleteVirtualNetwork.json
     */
    /**
     * Sample code: DeleteVirtualNetwork.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void deleteVirtualNetwork(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.virtualNetworks().delete("testrg", "ProdNetwork", null, Context.NONE);
    }
}
