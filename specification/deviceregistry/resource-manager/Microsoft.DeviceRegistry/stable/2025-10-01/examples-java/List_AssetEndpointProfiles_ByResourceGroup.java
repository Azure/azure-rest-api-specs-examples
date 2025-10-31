
/**
 * Samples for AssetEndpointProfiles ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/List_AssetEndpointProfiles_ByResourceGroup.json
     */
    /**
     * Sample code: List_AssetEndpointProfiles_ByResourceGroup.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void listAssetEndpointProfilesByResourceGroup(
        com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.assetEndpointProfiles().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
