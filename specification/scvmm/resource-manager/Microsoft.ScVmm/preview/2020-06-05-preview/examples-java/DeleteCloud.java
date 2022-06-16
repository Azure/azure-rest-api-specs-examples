import com.azure.core.util.Context;

/** Samples for Clouds Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/DeleteCloud.json
     */
    /**
     * Sample code: DeleteCloud.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void deleteCloud(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.clouds().delete("testrg", "HRCloud", null, Context.NONE);
    }
}
