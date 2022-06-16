import com.azure.core.util.Context;

/** Samples for VirtualNetworks GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/GetVirtualNetwork.json
     */
    /**
     * Sample code: GetVirtualNetwork.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void getVirtualNetwork(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualNetworks().getByResourceGroupWithResponse("testrg", "HRVirtualNetwork", Context.NONE);
    }
}
