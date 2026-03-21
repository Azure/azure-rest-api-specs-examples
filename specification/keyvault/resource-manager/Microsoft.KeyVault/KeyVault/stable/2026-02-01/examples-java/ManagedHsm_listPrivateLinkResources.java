
/**
 * Samples for MhsmPrivateLinkResources ListByMhsmResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ManagedHsm_listPrivateLinkResources.json
     */
    /**
     * Sample code: KeyVaultListPrivateLinkResources.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void keyVaultListPrivateLinkResources(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getMhsmPrivateLinkResources().listByMhsmResourceWithResponse("sample-group",
            "sample-mhsm", com.azure.core.util.Context.NONE);
    }
}
