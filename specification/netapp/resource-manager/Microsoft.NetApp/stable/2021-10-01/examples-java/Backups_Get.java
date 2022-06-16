import com.azure.core.util.Context;

/** Samples for Backups Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-10-01/examples/Backups_Get.json
     */
    /**
     * Sample code: Backups_Get.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupsGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.backups().getWithResponse("myRG", "account1", "pool1", "volume1", "backup1", Context.NONE);
    }
}
