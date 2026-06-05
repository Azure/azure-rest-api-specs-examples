
/**
 * Samples for Snapshots List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/SnapshotsList.json
     */
    /**
     * Sample code: List Snapshots.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void listSnapshots(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getSnapshots().list(com.azure.core.util.Context.NONE);
    }
}
