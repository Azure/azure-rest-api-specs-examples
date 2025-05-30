
/**
 * Samples for ReplicationRecoveryServicesProviders List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationRecoveryServicesProviders_List.json
     */
    /**
     * Sample code: Gets the list of registered recovery services providers in the vault. This is a view only api.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfRegisteredRecoveryServicesProvidersInTheVaultThisIsAViewOnlyApi(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationRecoveryServicesProviders().list("resourceGroupPS1", "vault1",
            com.azure.core.util.Context.NONE);
    }
}
