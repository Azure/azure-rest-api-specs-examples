
import com.azure.resourcemanager.appcontainers.models.CircuitBreakerPolicy;
import com.azure.resourcemanager.appcontainers.models.HeaderMatch;
import com.azure.resourcemanager.appcontainers.models.HttpConnectionPool;
import com.azure.resourcemanager.appcontainers.models.HttpRetryPolicy;
import com.azure.resourcemanager.appcontainers.models.TcpConnectionPool;
import com.azure.resourcemanager.appcontainers.models.TcpRetryPolicy;
import com.azure.resourcemanager.appcontainers.models.TimeoutPolicy;
import java.util.Arrays;

/**
 * Samples for AppResiliency CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/AppResiliency_CreateOrUpdate
     * .json
     */
    /**
     * Sample code: Create or Update App Resiliency.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        createOrUpdateAppResiliency(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.appResiliencies().define("resiliency-policy-1").withExistingContainerApp("rg", "testcontainerApp0")
            .withTimeoutPolicy(new TimeoutPolicy().withResponseTimeoutInSeconds(15).withConnectionTimeoutInSeconds(5))
            .withHttpRetryPolicy(new HttpRetryPolicy().withMaxRetries(5).withInitialDelayInMilliseconds(1000L)
                .withMaxIntervalInMilliseconds(10000L)
                .withHeaders(
                    Arrays.asList(new HeaderMatch().withHeaderProperty("X-Content-Type").withPrefixMatch("GOATS")))
                .withHttpStatusCodes(Arrays.asList(502, 503)).withErrors(
                    Arrays.asList("5xx", "connect-failure", "reset", "retriable-headers", "retriable-status-codes")))
            .withTcpRetryPolicy(new TcpRetryPolicy().withMaxConnectAttempts(3))
            .withCircuitBreakerPolicy(new CircuitBreakerPolicy().withConsecutiveErrors(5).withIntervalInSeconds(10)
                .withMaxEjectionPercent(50))
            .withHttpConnectionPool(
                new HttpConnectionPool().withHttp1MaxPendingRequests(1024).withHttp2MaxRequests(1024))
            .withTcpConnectionPool(new TcpConnectionPool().withMaxConnections(100)).create();
    }
}
