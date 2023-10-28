/** Samples for ManagedHsms GetDeleted. */
public final class Main {
    /*
     * x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2023-07-01/examples/DeletedManagedHsm_Get.json
     */
    /**
     * Sample code: Retrieve a deleted managed HSM.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void retrieveADeletedManagedHSM(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .vaults()
            .manager()
            .serviceClient()
            .getManagedHsms()
            .getDeletedWithResponse("hsm1", "westus", com.azure.core.util.Context.NONE);
    }
}
