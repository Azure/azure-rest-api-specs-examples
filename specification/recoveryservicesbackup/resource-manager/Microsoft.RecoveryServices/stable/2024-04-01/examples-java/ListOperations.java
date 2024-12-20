
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/
     * ListOperations.json
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
