
/**
 * Samples for Snapshots ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/SnapshotsListByResourceGroup.json
     */
    /**
     * Sample code: List Snapshots by Resource Group.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        listSnapshotsByResourceGroup(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getSnapshots().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
