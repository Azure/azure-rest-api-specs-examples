
/**
 * Samples for AssetEndpointProfiles GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/Get_AssetEndpointProfile.json
     */
    /**
     * Sample code: Get_AssetEndpointProfile.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void getAssetEndpointProfile(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.assetEndpointProfiles().getByResourceGroupWithResponse("myResourceGroup", "my-assetendpointprofile",
            com.azure.core.util.Context.NONE);
    }
}
