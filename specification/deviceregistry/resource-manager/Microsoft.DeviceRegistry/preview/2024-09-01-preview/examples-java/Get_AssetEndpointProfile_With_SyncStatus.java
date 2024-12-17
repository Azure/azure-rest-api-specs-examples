
/**
 * Samples for AssetEndpointProfiles GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/Get_AssetEndpointProfile_With_SyncStatus.json
     */
    /**
     * Sample code: Get_AssetEndpointProfile_With_SyncStatus.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        getAssetEndpointProfileWithSyncStatus(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.assetEndpointProfiles().getByResourceGroupWithResponse("myResourceGroup", "my-assetendpointprofile",
            com.azure.core.util.Context.NONE);
    }
}
