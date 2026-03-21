
/**
 * Samples for ManagedHsms ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ManagedHsm_ListByResourceGroup.json
     */
    /**
     * Sample code: List managed HSM Pools in a resource group.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void listManagedHSMPoolsInAResourceGroup(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getManagedHsms().listByResourceGroup("hsm-group", null,
            com.azure.core.util.Context.NONE);
    }
}
