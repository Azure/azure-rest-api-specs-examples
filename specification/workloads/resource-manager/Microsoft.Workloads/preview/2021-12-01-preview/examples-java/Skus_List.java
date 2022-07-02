import com.azure.core.util.Context;

/** Samples for Skus List. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/Skus_List.json
     */
    /**
     * Sample code: Skus.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void skus(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.skus().list(Context.NONE);
    }
}
