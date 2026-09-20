
import com.azure.resourcemanager.appcontainers.models.CheckNameAvailabilityRequest;

/**
 * Samples for Namespaces CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerApps_CheckNameAvailability.json
     */
    /**
     * Sample code: ContainerApps_CheckNameAvailability.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        containerAppsCheckNameAvailability(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.namespaces().checkNameAvailabilityWithResponse("examplerg", "testcontainerenv",
            new CheckNameAvailabilityRequest().withName("testcappname").withType("Microsoft.App/containerApps"),
            com.azure.core.util.Context.NONE);
    }
}
