import com.azure.core.util.Context;

/** Samples for Services ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/legacy/ServiceListByResourceGroup.json
     */
    /**
     * Sample code: List all services in resource group.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void listAllServicesInResourceGroup(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.services().listByResourceGroup("rgname", Context.NONE);
    }
}
