
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/ListOperations.json
     */
    /**
     * Sample code: ListOperations.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void listOperations(com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
