
/**
 * Samples for BackupInstancesExtensionRouting List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/
     * BackupInstanceOperations/ListBackupInstancesExtensionRouting.json
     */
    /**
     * Sample code: List BackupInstances associated with an azure resource.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void listBackupInstancesAssociatedWithAnAzureResource(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstancesExtensionRoutings().list(
            "subscriptions/36d32b25-3dc7-41b0-bde1-397500644591/resourceGroups/testRG/providers/Microsoft.Compute/disks/testDisk",
            com.azure.core.util.Context.NONE);
    }
}
