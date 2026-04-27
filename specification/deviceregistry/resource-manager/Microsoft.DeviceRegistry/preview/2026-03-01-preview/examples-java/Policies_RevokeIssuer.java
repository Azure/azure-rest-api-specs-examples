
/**
 * Samples for Policies RevokeIssuer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Policies_RevokeIssuer.json
     */
    /**
     * Sample code: Policies_RevokeIssuer.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void policiesRevokeIssuer(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.policies().revokeIssuer("rgdeviceregistry", "mynamespace", "mypolicy",
            com.azure.core.util.Context.NONE);
    }
}
