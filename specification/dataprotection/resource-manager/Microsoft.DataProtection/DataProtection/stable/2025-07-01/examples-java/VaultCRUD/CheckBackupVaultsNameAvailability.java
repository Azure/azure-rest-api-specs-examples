
import com.azure.resourcemanager.dataprotection.models.CheckNameAvailabilityRequest;

/**
 * Samples for BackupVaults CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VaultCRUD/CheckBackupVaultsNameAvailability.json
     */
    /**
     * Sample code: Check BackupVaults name availability.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        checkBackupVaultsNameAvailability(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupVaults()
            .checkNameAvailabilityWithResponse("SampleResourceGroup", "westus", new CheckNameAvailabilityRequest()
                .withName("swaggerExample").withType("Microsoft.DataProtection/BackupVaults"),
                com.azure.core.util.Context.NONE);
    }
}
