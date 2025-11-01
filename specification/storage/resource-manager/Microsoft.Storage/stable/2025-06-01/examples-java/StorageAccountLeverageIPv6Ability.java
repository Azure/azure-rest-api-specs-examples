
import com.azure.resourcemanager.storage.models.Action;
import com.azure.resourcemanager.storage.models.DefaultAction;
import com.azure.resourcemanager.storage.models.DualStackEndpointPreference;
import com.azure.resourcemanager.storage.models.IpRule;
import com.azure.resourcemanager.storage.models.NetworkRuleSet;
import com.azure.resourcemanager.storage.models.StorageAccountUpdateParameters;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for StorageAccounts Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2025-06-01/examples/
     * StorageAccountLeverageIPv6Ability.json
     */
    /**
     * Sample code: StorageAccountUpdateEnableIpv6Features.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountUpdateEnableIpv6Features(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getStorageAccounts().updateWithResponse("res9407", "sto8596",
            new StorageAccountUpdateParameters()
                .withNetworkRuleSet(new NetworkRuleSet()
                    .withIpv6Rules(Arrays
                        .asList(new IpRule().withIpAddressOrRange("2001:0db8:85a3::/64").withAction(Action.ALLOW)))
                    .withDefaultAction(DefaultAction.DENY))
                .withDualStackEndpointPreference(new DualStackEndpointPreference().withPublishIpv6Endpoint(true)),
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
