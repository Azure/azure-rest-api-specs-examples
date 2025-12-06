
/**
 * Samples for Subvolumes Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Subvolumes_Create.json
     */
    /**
     * Sample code: Subvolumes_Create.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void subvolumesCreate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.subvolumes().define("subvolume1").withExistingVolume("myRG", "account1", "pool1", "volume1")
            .withPath("/subvolumePath").create();
    }
}
