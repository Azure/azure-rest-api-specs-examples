import com.azure.resourcemanager.dnsresolver.models.ForwardingRuleState;
import com.azure.resourcemanager.dnsresolver.models.TargetDnsServer;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for ForwardingRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/ForwardingRule_Put.json
     */
    /**
     * Sample code: Upsert forwarding rule in a DNS forwarding ruleset.
     *
     * @param manager Entry point to DnsResolverManager.
     */
    public static void upsertForwardingRuleInADNSForwardingRuleset(
        com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager
            .forwardingRules()
            .define("sampleForwardingRule")
            .withExistingDnsForwardingRuleset("sampleResourceGroup", "sampleDnsForwardingRuleset")
            .withDomainName("contoso.com.")
            .withTargetDnsServers(
                Arrays
                    .asList(
                        new TargetDnsServer().withIpAddress("10.0.0.1").withPort(53),
                        new TargetDnsServer().withIpAddress("10.0.0.2").withPort(53)))
            .withMetadata(mapOf("additionalProp1", "value1"))
            .withForwardingRuleState(ForwardingRuleState.ENABLED)
            .create();
    }

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
