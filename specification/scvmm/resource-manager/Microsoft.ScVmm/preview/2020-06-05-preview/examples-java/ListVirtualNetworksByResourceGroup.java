import com.azure.core.util.Context;

/** Samples for VirtualNetworks ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/ListVirtualNetworksByResourceGroup.json
     */
    /**
     * Sample code: ListVirtualNetworksByResourceGroup.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void listVirtualNetworksByResourceGroup(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualNetworks().listByResourceGroup("testrg", Context.NONE);
    }
}
