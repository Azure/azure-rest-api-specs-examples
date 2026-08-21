
import com.azure.resourcemanager.storagesync.models.CloudEndpoint;
import com.azure.resourcemanager.storagesync.models.CloudEndpointUpdateProperties;

/**
 * Samples for CloudEndpoints Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/CloudEndpoints_Update.json
     */
    /**
     * Sample code: CloudEndpoints_Update.
     * 
     * @param manager Entry point to StorageSyncManager.
     */
    public static void cloudEndpointsUpdate(com.azure.resourcemanager.storagesync.StorageSyncManager manager) {
        CloudEndpoint resource = manager.cloudEndpoints().getWithResponse("rgstoragesync", "llg",
            "wwuoouzucgvfrsvjfgsobajg", "mjpalurfyrwkmqeygi", com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new CloudEndpointUpdateProperties().withChangeEnumerationIntervalDays(14))
            .apply();
    }
}
