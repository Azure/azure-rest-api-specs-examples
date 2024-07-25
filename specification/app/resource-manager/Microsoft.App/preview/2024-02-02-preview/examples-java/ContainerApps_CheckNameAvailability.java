
import com.azure.resourcemanager.appcontainers.models.CheckNameAvailabilityRequest;

/**
 * Samples for Namespaces CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * ContainerApps_CheckNameAvailability.json
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
