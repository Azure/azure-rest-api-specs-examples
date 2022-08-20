import com.azure.core.util.Context;

/** Samples for Clusters GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/GetCluster.json
     */
    /**
     * Sample code: GetCluster.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void getCluster(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.clusters().getByResourceGroupWithResponse("testrg", "HRCluster", Context.NONE);
    }
}
