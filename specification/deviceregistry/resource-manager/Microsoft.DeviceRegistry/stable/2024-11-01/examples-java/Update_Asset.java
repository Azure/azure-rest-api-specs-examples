
import com.azure.resourcemanager.deviceregistry.models.Asset;
import com.azure.resourcemanager.deviceregistry.models.AssetUpdateProperties;

/**
 * Samples for Assets Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/Update_Asset.json
     */
    /**
     * Sample code: Update_Asset.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void updateAsset(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        Asset resource = manager.assets()
            .getByResourceGroupWithResponse("myResourceGroup", "my-asset", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withProperties(new AssetUpdateProperties().withEnabled(false).withDisplayName("NewAssetDisplayName"))
            .apply();
    }
}
