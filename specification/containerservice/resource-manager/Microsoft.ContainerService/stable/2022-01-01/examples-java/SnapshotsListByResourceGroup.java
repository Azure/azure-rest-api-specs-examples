import com.azure.core.util.Context;

/** Samples for Snapshots ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-01-01/examples/SnapshotsListByResourceGroup.json
     */
    /**
     * Sample code: List Snapshots by Resource Group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSnapshotsByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getSnapshots().listByResourceGroup("rg1", Context.NONE);
    }
}
