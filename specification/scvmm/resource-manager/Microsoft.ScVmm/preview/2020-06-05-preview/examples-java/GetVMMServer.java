import com.azure.core.util.Context;

/** Samples for VmmServers GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/GetVMMServer.json
     */
    /**
     * Sample code: GetVMMServer.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void getVMMServer(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.vmmServers().getByResourceGroupWithResponse("testrg", "ContosoVMMServer", Context.NONE);
    }
}
