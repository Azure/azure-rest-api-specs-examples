
/**
 * Samples for Policies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Delete_Policies.json
     */
    /**
     * Sample code: Delete_Policies.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void deletePolicies(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.policies().delete("rgdeviceregistry", "mynamespace", "mypolicy", com.azure.core.util.Context.NONE);
    }
}
