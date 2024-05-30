
/**
 * Samples for Subvolumes Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-11-01/examples/Subvolumes_Create.json
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
