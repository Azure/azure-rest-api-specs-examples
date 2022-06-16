import com.azure.core.util.Context;

/** Samples for VmmServers ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/ListVMMServersByResourceGroup.json
     */
    /**
     * Sample code: ListVmmServersByResourceGroup.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void listVmmServersByResourceGroup(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.vmmServers().listByResourceGroup("testrg", Context.NONE);
    }
}
