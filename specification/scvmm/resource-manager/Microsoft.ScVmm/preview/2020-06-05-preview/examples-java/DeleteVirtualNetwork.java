import com.azure.core.util.Context;

/** Samples for VirtualNetworks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/DeleteVirtualNetwork.json
     */
    /**
     * Sample code: DeleteVirtualNetwork.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void deleteVirtualNetwork(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.virtualNetworks().delete("testrg", "HRVirtualNetwork", null, Context.NONE);
    }
}
