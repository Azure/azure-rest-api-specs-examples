/** Samples for Backups Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-11-01/examples/Backups_Create.json
     */
    /**
     * Sample code: Backups_Create.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupsCreate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager
            .backups()
            .define("backup1")
            .withRegion("eastus")
            .withExistingVolume("myRG", "account1", "pool1", "volume1")
            .withLabel("myLabel")
            .create();
    }
}
