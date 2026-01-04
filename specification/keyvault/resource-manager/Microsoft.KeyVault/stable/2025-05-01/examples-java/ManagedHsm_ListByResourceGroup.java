
/**
 * Samples for ManagedHsms ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ManagedHsm_ListByResourceGroup.json
     */
    /**
     * Sample code: List managed HSM Pools in a resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listManagedHSMPoolsInAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getManagedHsms().listByResourceGroup("hsm-group", null,
            com.azure.core.util.Context.NONE);
    }
}
