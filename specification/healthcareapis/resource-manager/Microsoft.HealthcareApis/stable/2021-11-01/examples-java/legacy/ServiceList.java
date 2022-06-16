import com.azure.core.util.Context;

/** Samples for Services List. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/legacy/ServiceList.json
     */
    /**
     * Sample code: List all services in subscription.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void listAllServicesInSubscription(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.services().list(Context.NONE);
    }
}
