import com.azure.core.util.Context;

/** Samples for Services GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/legacy/ServiceGet.json
     */
    /**
     * Sample code: Get metadata.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void getMetadata(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.services().getByResourceGroupWithResponse("rg1", "service1", Context.NONE);
    }
}
