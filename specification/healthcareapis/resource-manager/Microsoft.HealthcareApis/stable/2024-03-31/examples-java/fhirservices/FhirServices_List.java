
/**
 * Samples for FhirServices ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/fhirservices/
     * FhirServices_List.json
     */
    /**
     * Sample code: List fhirservices.
     * 
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void listFhirservices(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.fhirServices().listByWorkspace("testRG", "workspace1", com.azure.core.util.Context.NONE);
    }
}
