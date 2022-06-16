import com.azure.core.util.Context;

/** Samples for Clouds GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/GetCloud.json
     */
    /**
     * Sample code: GetCloud.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void getCloud(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.clouds().getByResourceGroupWithResponse("testrg", "HRCloud", Context.NONE);
    }
}
