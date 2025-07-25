
import com.azure.resourcemanager.networkcloud.models.CloudServicesNetworkEnableDefaultEgressEndpoints;
import com.azure.resourcemanager.networkcloud.models.EgressEndpoint;
import com.azure.resourcemanager.networkcloud.models.EndpointDependency;
import com.azure.resourcemanager.networkcloud.models.ExtendedLocation;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for CloudServicesNetworks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/
     * CloudServicesNetworks_Create.json
     */
    /**
     * Sample code: Create or update cloud services network.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        createOrUpdateCloudServicesNetwork(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.cloudServicesNetworks().define("cloudServicesNetworkName").withRegion("location")
            .withExistingResourceGroup("resourceGroupName")
            .withExtendedLocation(new ExtendedLocation().withName(
                "/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName")
                .withType("CustomLocation"))
            .withTags(mapOf("key1", "fakeTokenPlaceholder", "key2", "fakeTokenPlaceholder"))
            .withAdditionalEgressEndpoints(Arrays.asList(new EgressEndpoint().withCategory("azure-resource-management")
                .withEndpoints(Arrays.asList(
                    new EndpointDependency().withDomainName("storageaccountex.blob.core.windows.net").withPort(443L)))))
            .withEnableDefaultEgressEndpoints(CloudServicesNetworkEnableDefaultEgressEndpoints.FALSE).create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
