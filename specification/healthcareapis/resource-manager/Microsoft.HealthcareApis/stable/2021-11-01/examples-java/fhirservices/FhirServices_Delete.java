import com.azure.core.util.Context;

/** Samples for FhirServices Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/fhirservices/FhirServices_Delete.json
     */
    /**
     * Sample code: Delete a Fhir Service.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void deleteAFhirService(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.fhirServices().delete("testRG", "fhirservice1", "workspace1", Context.NONE);
    }
}
