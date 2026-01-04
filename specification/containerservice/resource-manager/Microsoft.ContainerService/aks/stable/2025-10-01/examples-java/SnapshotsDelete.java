
/**
 * Samples for Snapshots Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * SnapshotsDelete.json
     */
    /**
     * Sample code: Delete Snapshot.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteSnapshot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getSnapshots().deleteWithResponse("rg1", "snapshot1",
            com.azure.core.util.Context.NONE);
    }
}
