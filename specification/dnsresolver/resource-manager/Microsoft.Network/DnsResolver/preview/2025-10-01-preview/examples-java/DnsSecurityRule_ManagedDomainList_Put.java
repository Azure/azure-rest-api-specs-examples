
import com.azure.resourcemanager.dnsresolver.models.ActionType;
import com.azure.resourcemanager.dnsresolver.models.DnsSecurityRuleAction;
import com.azure.resourcemanager.dnsresolver.models.DnsSecurityRuleState;
import com.azure.resourcemanager.dnsresolver.models.ManagedDomainList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DnsSecurityRules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/DnsSecurityRule_ManagedDomainList_Put.json
     */
    /**
     * Sample code: Upsert DNS security rule with managed domain list.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        upsertDNSSecurityRuleWithManagedDomainList(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsSecurityRules().define("sampleDnsSecurityRule").withRegion("westus2")
            .withExistingDnsResolverPolicy("sampleResourceGroup", "sampleDnsResolverPolicy").withPriority(100)
            .withAction(new DnsSecurityRuleAction().withActionType(ActionType.BLOCK))
            .withTags(mapOf("key1", "fakeTokenPlaceholder"))
            .withManagedDomainLists(Arrays.asList(ManagedDomainList.AZURE_DNS_THREAT_INTEL))
            .withDnsSecurityRuleState(DnsSecurityRuleState.ENABLED).create();
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
