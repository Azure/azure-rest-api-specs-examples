import com.azure.core.util.Context;

/** Samples for DicomServices ListByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/dicomservices/DicomServices_List.json
     */
    /**
     * Sample code: List dicomservices.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void listDicomservices(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.dicomServices().listByWorkspace("testRG", "workspace1", Context.NONE);
    }
}
