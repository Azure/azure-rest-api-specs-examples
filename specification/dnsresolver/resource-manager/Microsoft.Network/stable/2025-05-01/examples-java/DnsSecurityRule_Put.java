
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.dnsresolver.models.ActionType;
import com.azure.resourcemanager.dnsresolver.models.DnsSecurityRuleAction;
import com.azure.resourcemanager.dnsresolver.models.DnsSecurityRuleState;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DnsSecurityRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/DnsSecurityRule_Put.json
     */
    /**
     * Sample code: Upsert DNS security rule.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void upsertDNSSecurityRule(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsSecurityRules().define("sampleDnsSecurityRule").withRegion("westus2")
            .withExistingDnsResolverPolicy("sampleResourceGroup", "sampleDnsResolverPolicy").withPriority(100)
            .withAction(new DnsSecurityRuleAction().withActionType(ActionType.BLOCK))
            .withDnsResolverDomainLists(Arrays.asList(new SubResource().withId(
                "/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsResolverDomainLists/sampleDnsResolverDomainList")))
            .withTags(mapOf("key1", "fakeTokenPlaceholder")).withDnsSecurityRuleState(DnsSecurityRuleState.ENABLED)
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
