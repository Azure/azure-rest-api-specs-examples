
import com.azure.resourcemanager.deviceregistry.models.AuthenticationMethod;
import com.azure.resourcemanager.deviceregistry.models.DiscoveredAssetEndpointProfileProperties;
import com.azure.resourcemanager.deviceregistry.models.ExtendedLocation;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DiscoveredAssetEndpointProfiles CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01-preview/Create_DiscoveredAssetEndpointProfile.json
     */
    /**
     * Sample code: Create_DiscoveredAssetEndpointProfile.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        createDiscoveredAssetEndpointProfile(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.discoveredAssetEndpointProfiles().define("my-discoveredassetendpointprofile").withRegion("West Europe")
            .withExistingResourceGroup("myResourceGroup")
            .withExtendedLocation(new ExtendedLocation().withType("CustomLocation").withName(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"))
            .withTags(mapOf("site", "building-1"))
            .withProperties(new DiscoveredAssetEndpointProfileProperties()
                .withTargetAddress("https://www.example.com/myTargetAddress")
                .withAdditionalConfiguration("{\"foo\": \"bar\"}")
                .withSupportedAuthenticationMethods(Arrays.asList(AuthenticationMethod.ANONYMOUS,
                    AuthenticationMethod.CERTIFICATE, AuthenticationMethod.USERNAME_PASSWORD))
                .withEndpointProfileType("myEndpointProfileType")
                .withDiscoveryId("11111111-1111-1111-1111-111111111111").withVersion(73766L))
            .create();
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
