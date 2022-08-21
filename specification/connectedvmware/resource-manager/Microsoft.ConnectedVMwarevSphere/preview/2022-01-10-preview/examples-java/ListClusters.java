import com.azure.core.util.Context;

/** Samples for Clusters List. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/ListClusters.json
     */
    /**
     * Sample code: ListClusters.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listClusters(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.clusters().list(Context.NONE);
    }
}
