
import com.azure.resourcemanager.deviceregistry.models.AssetEndpointProfileProperties;
import com.azure.resourcemanager.deviceregistry.models.Authentication;
import com.azure.resourcemanager.deviceregistry.models.AuthenticationMethod;
import com.azure.resourcemanager.deviceregistry.models.ExtendedLocation;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AssetEndpointProfiles CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/Create_AssetEndpointProfile.json
     */
    /**
     * Sample code: Create_AssetEndpointProfile.
     * 
     * @param manager Entry point to DeviceRegistryManager.
     */
    public static void
        createAssetEndpointProfile(com.azure.resourcemanager.deviceregistry.DeviceRegistryManager manager) {
        manager.assetEndpointProfiles().define("my-assetendpointprofile").withRegion("West Europe")
            .withExistingResourceGroup("myResourceGroup")
            .withExtendedLocation(new ExtendedLocation().withType("CustomLocation").withName(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1"))
            .withTags(mapOf("site", "building-1"))
            .withProperties(
                new AssetEndpointProfileProperties().withTargetAddress("https://www.example.com/myTargetAddress")
                    .withEndpointProfileType("myEndpointProfileType")
                    .withAuthentication(new Authentication().withMethod(AuthenticationMethod.ANONYMOUS)))
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
