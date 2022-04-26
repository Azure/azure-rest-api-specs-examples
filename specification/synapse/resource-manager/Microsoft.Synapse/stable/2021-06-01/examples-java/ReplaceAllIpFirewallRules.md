Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.fluent.models.IpFirewallRuleProperties;
import com.azure.resourcemanager.synapse.models.ReplaceAllIpFirewallRulesRequest;
import java.util.HashMap;
import java.util.Map;

/** Samples for IpFirewallRules ReplaceAll. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ReplaceAllIpFirewallRules.json
     */
    /**
     * Sample code: Replace all IP firewall rules in a workspace.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void replaceAllIPFirewallRulesInAWorkspace(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .ipFirewallRules()
            .replaceAll(
                "ExampleResourceGroup",
                "ExampleWorkspace",
                new ReplaceAllIpFirewallRulesRequest()
                    .withIpFirewallRules(
                        mapOf(
                            "AnotherExampleFirewallRule",
                            new IpFirewallRuleProperties()
                                .withEndIpAddress("10.0.1.254")
                                .withStartIpAddress("10.0.1.0"),
                            "ExampleFirewallRule",
                            new IpFirewallRuleProperties()
                                .withEndIpAddress("10.0.0.254")
                                .withStartIpAddress("10.0.0.0"))),
                Context.NONE);
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
```
