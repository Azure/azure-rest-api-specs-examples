
import com.azure.resourcemanager.newrelicobservability.models.ActivateSaaSParameterRequest;

/**
 * Samples for SaaS ActivateResource.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/newrelic/resource-manager/NewRelic.Observability/preview/2025-05-01-preview/examples/ActivateSaaS.
     * json
     */
    /**
     * Sample code: ActivateSaaS.
     * 
     * @param manager Entry point to NewRelicObservabilityManager.
     */
    public static void
        activateSaaS(com.azure.resourcemanager.newrelicobservability.NewRelicObservabilityManager manager) {
        manager.saaS()
            .activateResourceWithResponse(new ActivateSaaSParameterRequest()
                .withSaasGuid("00000000-0000-0000-0000-000005430000").withPublisherId("publisherId"),
                com.azure.core.util.Context.NONE);
    }
}
