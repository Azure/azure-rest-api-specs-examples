
/**
 * Samples for ManagedHsms ListDeleted.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/DeletedManagedHsm_List.json
     */
    /**
     * Sample code: List deleted managed HSMs in the specified subscription.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void
        listDeletedManagedHSMsInTheSpecifiedSubscription(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getManagedHsms().listDeleted(com.azure.core.util.Context.NONE);
    }
}
