
/**
 * Samples for LongTermRetentionBackups RemoveTimeBasedImmutability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/RemoveTimeBasedImmutabilityLongTermRetentionBackup.json
     */
    /**
     * Sample code: Remove time based immutability of the long term retention backup.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void removeTimeBasedImmutabilityOfTheLongTermRetentionBackup(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getLongTermRetentionBackups().removeTimeBasedImmutability("japaneast", "testserver",
            "testDatabase", "55555555-6666-7777-8888-999999999999;131637960820000000;Hot",
            com.azure.core.util.Context.NONE);
    }
}
