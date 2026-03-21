
/**
 * Samples for ManagedHsms GetDeleted.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/DeletedManagedHsm_Get.json
     */
    /**
     * Sample code: Retrieve a deleted managed HSM.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void retrieveADeletedManagedHSM(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getManagedHsms().getDeletedWithResponse("hsm1", "westus",
            com.azure.core.util.Context.NONE);
    }
}
