import com.azure.core.util.Context;

/** Samples for BackupPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-05-01/examples/BackupPolicies_Get.json
     */
    /**
     * Sample code: Backups_Get.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupsGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.backupPolicies().getWithResponse("myRG", "account1", "backupPolicyName", Context.NONE);
    }
}
