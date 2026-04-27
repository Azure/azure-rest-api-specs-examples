
/**
 * Samples for Policies ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/List_Policies_ByResourceGroup.json
     */
    /**
     * Sample code: List_Policies_ByResourceGroup.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        listPoliciesByResourceGroup(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.policies().listByResourceGroup("rgdeviceregistry", "mynamespace", com.azure.core.util.Context.NONE);
    }
}
