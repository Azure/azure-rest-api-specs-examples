
/**
 * Samples for Operations ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /Operations_List.json
     */
    /**
     * Sample code: Returns the list of available operations.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void returnsTheListOfAvailableOperations(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.operations().listByResourceGroup("resourceGroupPS1", com.azure.core.util.Context.NONE);
    }
}
