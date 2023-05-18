/** Samples for Vaults List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/ListBySubscriptionIds.json
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
