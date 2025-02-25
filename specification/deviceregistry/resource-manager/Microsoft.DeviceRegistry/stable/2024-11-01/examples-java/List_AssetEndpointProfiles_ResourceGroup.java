
/**
 * Samples for AssetEndpointProfiles ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/List_AssetEndpointProfiles_ResourceGroup.json
     */
    /**
     * Sample code: List_AssetEndpointProfiles_ResourceGroup.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        listAssetEndpointProfilesResourceGroup(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.assetEndpointProfiles().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
