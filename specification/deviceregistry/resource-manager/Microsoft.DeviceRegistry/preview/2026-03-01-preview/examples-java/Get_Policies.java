
/**
 * Samples for Policies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Get_Policies.json
     */
    /**
     * Sample code: Get_Policies.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void getPolicies(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.policies().getWithResponse("rgdeviceregistry", "mynamespace", "myPolicy",
            com.azure.core.util.Context.NONE);
    }
}
