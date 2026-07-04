
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.NetworkVirtualApplianceInner;
import com.azure.resourcemanager.network.models.DelegationProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for NetworkVirtualAppliances CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NetworkVirtualApplianceSaaSPut.json
     */
    /**
     * Sample code: Create SaaS NetworkVirtualAppliance.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createSaaSNetworkVirtualAppliance(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkVirtualAppliances().createOrUpdate("rg1", "nva",
            new NetworkVirtualApplianceInner().withLocation("West US").withTags(mapOf("key1", "fakeTokenPlaceholder"))
                .withVirtualHub(new SubResource().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1"))
                .withDelegation(new DelegationProperties().withServiceName("PaloAltoNetworks.Cloudngfw/firewalls")),
            com.azure.core.util.Context.NONE);
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
