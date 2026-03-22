
/**
 * Samples for ManagedHsms Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ManagedHsm_Delete.json
     */
    /**
     * Sample code: Delete a managed HSM Pool.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void deleteAManagedHSMPool(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getManagedHsms().delete("hsm-group", "hsm1", com.azure.core.util.Context.NONE);
    }
}
