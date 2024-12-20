
import com.azure.resourcemanager.notificationhubs.models.AccessRights;
import com.azure.resourcemanager.notificationhubs.models.IpRule;
import com.azure.resourcemanager.notificationhubs.models.NetworkAcls;
import com.azure.resourcemanager.notificationhubs.models.PublicInternetAuthorizationRule;
import com.azure.resourcemanager.notificationhubs.models.Sku;
import com.azure.resourcemanager.notificationhubs.models.SkuName;
import com.azure.resourcemanager.notificationhubs.models.ZoneRedundancyPreference;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Namespaces CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * Namespaces/CreateOrUpdate.json
     */
    /**
     * Sample code: Namespaces_CreateOrUpdate.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void
        namespacesCreateOrUpdate(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.namespaces().define("nh-sdk-ns").withRegion("South Central US").withExistingResourceGroup("5ktrial")
            .withSku(new Sku().withName(SkuName.STANDARD).withTier("Standard"))
            .withTags(mapOf("tag1", "value1", "tag2", "value2")).withZoneRedundancy(ZoneRedundancyPreference.ENABLED)
            .withNetworkAcls(new NetworkAcls()
                .withIpRules(Arrays.asList(new IpRule().withIpMask("185.48.100.00/24")
                    .withRights(Arrays.asList(AccessRights.MANAGE, AccessRights.SEND, AccessRights.LISTEN))))
                .withPublicNetworkRule(
                    new PublicInternetAuthorizationRule().withRights(Arrays.asList(AccessRights.LISTEN))))
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
