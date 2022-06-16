import com.azure.core.util.Context;

/** Samples for Clouds ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/ListCloudsByResourceGroup.json
     */
    /**
     * Sample code: ListCloudsByResourceGroup.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void listCloudsByResourceGroup(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.clouds().listByResourceGroup("testrg", Context.NONE);
    }
}
