
/**
 * Samples for DeletedVaults ListBySubscriptionId.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/DeletedVaults_ListBySubscriptionId.json
     */
    /**
     * Sample code: List deleted vaults in a subscription.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void
        listDeletedVaultsInASubscription(com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.deletedVaults().listBySubscriptionId("westus", com.azure.core.util.Context.NONE);
    }
}
