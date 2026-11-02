
import com.azure.resourcemanager.privatetrafficmanager.models.AdministrativeStatus;
import com.azure.resourcemanager.privatetrafficmanager.models.CustomTopologyMapMode;
import com.azure.resourcemanager.privatetrafficmanager.models.ProfileEndpoint;
import com.azure.resourcemanager.privatetrafficmanager.models.ProfileProperties;
import com.azure.resourcemanager.privatetrafficmanager.models.ProfileStatus;
import com.azure.resourcemanager.privatetrafficmanager.models.TrafficRoutingMethod;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Profiles CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Profiles_CreateOrUpdate_VNetAffinity.json
     */
    /**
     * Sample code: Profiles_CreateOrUpdate_VNetAffinity.
     * 
     * @param manager Entry point to PrivateTrafficManagerManager.
     */
    public static void profilesCreateOrUpdateVNetAffinity(
        com.azure.resourcemanager.privatetrafficmanager.PrivateTrafficManagerManager manager) {
        manager.profiles().define("vnet-affinity-profile").withRegion("global")
            .withExistingResourceGroup("rgprivateTrafficManager")
            .withProperties(new ProfileProperties().withCustomTopologyMapMode(CustomTopologyMapMode.DISABLED)
                .withProfileStatus(ProfileStatus.ENABLED).withTrafficRoutingMethod(TrafficRoutingMethod.VNET)
                .withEndpoints(Arrays.asList(new ProfileEndpoint().withName("matched-endpoint").withTarget("10.0.0.10")
                    .withEndpointStatus(AdministrativeStatus.ENABLED)
                    .withVirtualNetworkIds(Arrays.asList(
                        "/subscriptions/10B6D88D-ADF4-4281-B3D0-B5A6702DEEDA/resourceGroups/rgprivateTrafficManager/providers/Microsoft.Network/virtualNetworks/vnet-east")),
                    new ProfileEndpoint().withName("fallback-endpoint").withTarget("10.0.0.20")
                        .withEndpointStatus(AdministrativeStatus.ENABLED))))
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
