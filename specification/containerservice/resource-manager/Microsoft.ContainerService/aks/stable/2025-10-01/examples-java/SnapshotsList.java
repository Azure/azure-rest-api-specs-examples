
/**
 * Samples for Snapshots List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * SnapshotsList.json
     */
    /**
     * Sample code: List Snapshots.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSnapshots(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getSnapshots().list(com.azure.core.util.Context.NONE);
    }
}
