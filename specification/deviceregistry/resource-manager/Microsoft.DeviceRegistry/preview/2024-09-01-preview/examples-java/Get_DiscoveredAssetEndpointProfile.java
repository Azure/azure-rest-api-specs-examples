
/**
 * Samples for DiscoveredAssetEndpointProfiles GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/Get_DiscoveredAssetEndpointProfile.json
     */
    /**
     * Sample code: Get_DiscoveredAssetEndpointProfile.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        getDiscoveredAssetEndpointProfile(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.discoveredAssetEndpointProfiles().getByResourceGroupWithResponse("myResourceGroup",
            "my-discoveredassetendpointprofile", com.azure.core.util.Context.NONE);
    }
}
