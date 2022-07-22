import com.azure.core.util.Context;
import com.azure.resourcemanager.netapp.models.Backup;

/** Samples for Backups Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-03-01/examples/Backups_Update.json
     */
    /**
     * Sample code: Backups_Update.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupsUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        Backup resource =
            manager
                .backups()
                .getWithResponse("myRG", "account1", "pool1", "volume1", "backup1", Context.NONE)
                .getValue();
        resource.update().apply();
    }
}
