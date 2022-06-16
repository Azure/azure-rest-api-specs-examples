import com.azure.core.util.Context;

/** Samples for PrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/legacy/PrivateLinkResourceGet.json
     */
    /**
     * Sample code: PrivateLinkResources_Get.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void privateLinkResourcesGet(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.privateLinkResources().getWithResponse("rgname", "service1", "fhir", Context.NONE);
    }
}
