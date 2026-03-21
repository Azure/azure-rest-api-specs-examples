
/**
 * Samples for MhsmPrivateEndpointConnections ListByResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ManagedHsm_ListPrivateEndpointConnectionsByResource.json
     */
    /**
     * Sample code: List managed HSM Pools in a subscription.
     * 
     * @param manager Entry point to KeyVaultManager.
     */
    public static void listManagedHSMPoolsInASubscription(com.azure.resourcemanager.keyvault.KeyVaultManager manager) {
        manager.serviceClient().getMhsmPrivateEndpointConnections().listByResource("sample-group", "sample-mhsm",
            com.azure.core.util.Context.NONE);
    }
}
