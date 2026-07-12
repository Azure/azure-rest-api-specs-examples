
/**
 * Samples for Vaults List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/ListBySubscriptionIds.json
     */
    /**
     * Sample code: List of Recovery Services Resources in SubscriptionId.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void listOfRecoveryServicesResourcesInSubscriptionId(
        com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.vaults().list(com.azure.core.util.Context.NONE);
    }
}
