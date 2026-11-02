
import com.azure.resourcemanager.privatetrafficmanager.models.HealthPolicyProperties;
import com.azure.resourcemanager.privatetrafficmanager.models.ProbeConfig;
import com.azure.resourcemanager.privatetrafficmanager.models.ProbeHealthPolicy;
import com.azure.resourcemanager.privatetrafficmanager.models.Protocol;

/**
 * Samples for HealthPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/HealthPolicies_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: HealthPolicies_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to PrivateTrafficManagerManager.
     */
    public static void healthPoliciesCreateOrUpdateMaximumSet(
        com.azure.resourcemanager.privatetrafficmanager.PrivateTrafficManagerManager manager) {
        manager.healthPolicies().createOrUpdate("rgprivateTrafficManager", "myProfile", "myHealthPolicy",
            new ProbeHealthPolicy().withProperties(new HealthPolicyProperties()
                .withProbeConfig(new ProbeConfig().withProtocol(Protocol.HTTPS).withPort(443L).withPath("/health")
                    .withIntervalInSeconds(30L).withTimeoutInSeconds(10L).withToleratedNumberOfFailures(3L))),
            com.azure.core.util.Context.NONE);
    }
}
