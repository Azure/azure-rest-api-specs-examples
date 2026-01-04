
/**
 * Samples for MhsmPrivateEndpointConnections ListByResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ManagedHsm_ListPrivateEndpointConnectionsByResource.json
     */
    /**
     * Sample code: List managed HSM Pools in a subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listManagedHSMPoolsInASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getMhsmPrivateEndpointConnections().listByResource("sample-group",
            "sample-mhsm", com.azure.core.util.Context.NONE);
    }
}
