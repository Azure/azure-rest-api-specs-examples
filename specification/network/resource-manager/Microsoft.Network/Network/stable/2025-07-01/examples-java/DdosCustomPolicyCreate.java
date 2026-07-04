
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
     * x-ms-original-file: 2025-07-01/DdosCustomPolicyCreate.json
     */
    /**
     * Sample code: Create DDoS custom policy.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createDDoSCustomPolicy(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getDdosCustomPolicies()
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
