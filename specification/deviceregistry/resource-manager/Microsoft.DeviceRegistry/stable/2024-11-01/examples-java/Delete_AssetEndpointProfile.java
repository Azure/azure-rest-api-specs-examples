
/**
 * Samples for AssetEndpointProfiles Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/Delete_AssetEndpointProfile.json
     */
    /**
     * Sample code: Delete_AssetEndpointProfile.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        deleteAssetEndpointProfile(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.assetEndpointProfiles().delete("myResourceGroup", "my-assetendpointprofile",
            com.azure.core.util.Context.NONE);
    }
}
