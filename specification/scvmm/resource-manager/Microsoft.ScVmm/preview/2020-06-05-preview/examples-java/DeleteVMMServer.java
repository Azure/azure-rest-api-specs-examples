import com.azure.core.util.Context;

/** Samples for VmmServers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/DeleteVMMServer.json
     */
    /**
     * Sample code: DeleteVMMServer.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void deleteVMMServer(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.vmmServers().delete("testrg", "ContosoVMMServer", null, Context.NONE);
    }
}
