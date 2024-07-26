
import com.azure.resourcemanager.appcontainers.models.DaprComponentResiliencyPolicyCircuitBreakerPolicyConfiguration;
import com.azure.resourcemanager.appcontainers.models.DaprComponentResiliencyPolicyConfiguration;
import com.azure.resourcemanager.appcontainers.models.DaprComponentResiliencyPolicyHttpRetryBackOffConfiguration;
import com.azure.resourcemanager.appcontainers.models.DaprComponentResiliencyPolicyHttpRetryPolicyConfiguration;
import com.azure.resourcemanager.appcontainers.models.DaprComponentResiliencyPolicyTimeoutPolicyConfiguration;

/**
 * Samples for DaprComponentResiliencyPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * DaprComponentResiliencyPolicy_CreateOrUpdate_OutboundOnly.json
     */
    /**
     * Sample code: Create or update dapr component resiliency policy with outbound policy only.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createOrUpdateDaprComponentResiliencyPolicyWithOutboundPolicyOnly(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.daprComponentResiliencyPolicies().define("myresiliencypolicy")
            .withExistingDaprComponent("examplerg", "myenvironment", "mydaprcomponent")
            .withOutboundPolicy(new DaprComponentResiliencyPolicyConfiguration()
                .withHttpRetryPolicy(new DaprComponentResiliencyPolicyHttpRetryPolicyConfiguration().withMaxRetries(5)
                    .withRetryBackOff(new DaprComponentResiliencyPolicyHttpRetryBackOffConfiguration()
                        .withInitialDelayInMilliseconds(100).withMaxIntervalInMilliseconds(30000)))
                .withTimeoutPolicy(
                    new DaprComponentResiliencyPolicyTimeoutPolicyConfiguration().withResponseTimeoutInSeconds(12))
                .withCircuitBreakerPolicy(new DaprComponentResiliencyPolicyCircuitBreakerPolicyConfiguration()
                    .withConsecutiveErrors(3).withTimeoutInSeconds(20).withIntervalInSeconds(60)))
            .create();
    }
}
