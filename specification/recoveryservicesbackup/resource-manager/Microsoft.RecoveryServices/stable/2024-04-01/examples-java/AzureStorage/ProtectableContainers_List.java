
/**
 * Samples for ProtectableContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/
     * AzureStorage/ProtectableContainers_List.json
     */
    /**
     * Sample code: List protectable items with backupManagementType filter as AzureStorage.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void listProtectableItemsWithBackupManagementTypeFilterAsAzureStorage(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.protectableContainers().list("testvault", "testRg", "Azure",
            "backupManagementType eq 'AzureStorage' and workloadType eq 'AzureFileShare'",
            com.azure.core.util.Context.NONE);
    }
}
