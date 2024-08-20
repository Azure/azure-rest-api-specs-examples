
/**
 * Samples for Snapshots Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-03-01/examples/Snapshots_Create.json
     */
    /**
     * Sample code: Snapshots_Create.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void snapshotsCreate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.snapshots().define("snapshot1").withRegion("eastus")
            .withExistingVolume("myRG", "account1", "pool1", "volume1").create();
    }
}
