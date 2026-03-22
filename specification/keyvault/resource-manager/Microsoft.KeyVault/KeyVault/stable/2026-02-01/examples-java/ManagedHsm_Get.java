
/**
 * Samples for ManagedHsms GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ManagedHsm_Get.json
     */
    /**
     * Sample code: Retrieve a managed HSM Pool.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void retrieveAManagedHSMPool(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getManagedHsms().getByResourceGroupWithResponse("hsm-group", "hsm1",
            com.azure.core.util.Context.NONE);
    }
}
