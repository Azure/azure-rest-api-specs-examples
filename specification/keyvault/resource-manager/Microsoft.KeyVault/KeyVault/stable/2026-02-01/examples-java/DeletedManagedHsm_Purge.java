
/**
 * Samples for ManagedHsms PurgeDeleted.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/DeletedManagedHsm_Purge.json
     */
    /**
     * Sample code: Purge a managed HSM Pool.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void purgeAManagedHSMPool(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getManagedHsms().purgeDeleted("hsm1", "westus", com.azure.core.util.Context.NONE);
    }
}
