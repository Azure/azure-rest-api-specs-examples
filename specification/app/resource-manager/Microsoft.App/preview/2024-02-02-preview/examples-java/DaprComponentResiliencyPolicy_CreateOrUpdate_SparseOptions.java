
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
     * DaprComponentResiliencyPolicy_CreateOrUpdate_SparseOptions.json
     */
    /**
     * Sample code: Create or update dapr component resiliency policy with sparse options.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createOrUpdateDaprComponentResiliencyPolicyWithSparseOptions(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.daprComponentResiliencyPolicies().define("myresiliencypolicy")
            .withExistingDaprComponent("examplerg", "myenvironment", "mydaprcomponent")
            .withInboundPolicy(new DaprComponentResiliencyPolicyConfiguration()
                .withHttpRetryPolicy(new DaprComponentResiliencyPolicyHttpRetryPolicyConfiguration().withMaxRetries(5)
                    .withRetryBackOff(new DaprComponentResiliencyPolicyHttpRetryBackOffConfiguration()
                        .withInitialDelayInMilliseconds(2000).withMaxIntervalInMilliseconds(5500)))
                .withCircuitBreakerPolicy(new DaprComponentResiliencyPolicyCircuitBreakerPolicyConfiguration()
                    .withConsecutiveErrors(3).withTimeoutInSeconds(20)))
            .withOutboundPolicy(new DaprComponentResiliencyPolicyConfiguration().withTimeoutPolicy(
                new DaprComponentResiliencyPolicyTimeoutPolicyConfiguration().withResponseTimeoutInSeconds(12)))
            .create();
    }
}
