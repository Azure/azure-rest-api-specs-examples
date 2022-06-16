import com.azure.core.util.Context;

/** Samples for Clouds List. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/ListCloudsBySubscription.json
     */
    /**
     * Sample code: ListCloudsBySubscription.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void listCloudsBySubscription(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager.clouds().list(Context.NONE);
    }
}
