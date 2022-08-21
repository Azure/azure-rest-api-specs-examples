import com.azure.core.util.Context;

/** Samples for Clusters Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/DeleteCluster.json
     */
    /**
     * Sample code: DeleteCluster.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void deleteCluster(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.clusters().delete("testrg", "HRCluster", null, Context.NONE);
    }
}
