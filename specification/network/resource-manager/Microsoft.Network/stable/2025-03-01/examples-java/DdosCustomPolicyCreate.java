
import com.azure.resourcemanager.network.fluent.models.DdosCustomPolicyInner;
import com.azure.resourcemanager.network.models.DdosDetectionMode;
import com.azure.resourcemanager.network.models.DdosDetectionRule;
import com.azure.resourcemanager.network.models.DdosTrafficType;
import com.azure.resourcemanager.network.models.TrafficDetectionRule;
import java.util.Arrays;

/**
 * Samples for DdosCustomPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/DdosCustomPolicyCreate.json
     */
    /**
     * Sample code: Create DDoS custom policy.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createDDoSCustomPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getDdosCustomPolicies()
            .createOrUpdate("rg1", "test-ddos-custom-policy",
                new DdosCustomPolicyInner().withLocation("centraluseuap")
                    .withDetectionRules(
                        Arrays
                            .asList(new DdosDetectionRule().withName("detectionRuleTcp")
                                .withDetectionMode(DdosDetectionMode.TRAFFIC_THRESHOLD)
                                .withTrafficDetectionRule(new TrafficDetectionRule()
                                    .withTrafficType(DdosTrafficType.TCP).withPacketsPerSecond(1000000)))),
                com.azure.core.util.Context.NONE);
    }
}
