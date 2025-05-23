
import com.azure.resourcemanager.deviceregistry.models.DiscoveredAsset;
import com.azure.resourcemanager.deviceregistry.models.DiscoveredAssetUpdateProperties;
import com.azure.resourcemanager.deviceregistry.models.Topic;
import com.azure.resourcemanager.deviceregistry.models.TopicRetainType;

/**
 * Samples for DiscoveredAssets Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/Update_DiscoveredAsset.json
     */
    /**
     * Sample code: Update_DiscoveredAsset.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void updateDiscoveredAsset(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        DiscoveredAsset resource = manager.discoveredAssets()
            .getByResourceGroupWithResponse("myResourceGroup", "my-discoveredasset", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update()
            .withProperties(
                new DiscoveredAssetUpdateProperties().withDocumentationUri("https://www.example.com/manual-2")
                    .withDefaultTopic(new Topic().withPath("/path/defaultTopic").withRetain(TopicRetainType.NEVER)))
            .apply();
    }
}
