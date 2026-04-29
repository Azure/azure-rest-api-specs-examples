
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/ListOperations.json
     */
    /**
     * Sample code: ListOperations.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void
        listOperations(com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
