
/**
 * Samples for ManagedHsms GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ManagedHsm_Get.json
     */
    /**
     * Sample code: Retrieve a managed HSM Pool.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void retrieveAManagedHSMPool(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getManagedHsms().getByResourceGroupWithResponse("hsm-group", "hsm1",
            com.azure.core.util.Context.NONE);
    }
}
