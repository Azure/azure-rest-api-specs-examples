
/**
 * Samples for ManagedHsms Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ManagedHsm_Delete.json
     */
    /**
     * Sample code: Delete a managed HSM Pool.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAManagedHSMPool(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getManagedHsms().delete("hsm-group", "hsm1",
            com.azure.core.util.Context.NONE);
    }
}
