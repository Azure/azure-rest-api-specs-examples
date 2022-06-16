import com.azure.core.util.Context;

/** Samples for DicomServices Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/dicomservices/DicomServices_Delete.json
     */
    /**
     * Sample code: Delete a dicomservice.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void deleteADicomservice(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.dicomServices().delete("testRG", "blue", "workspace1", Context.NONE);
    }
}
