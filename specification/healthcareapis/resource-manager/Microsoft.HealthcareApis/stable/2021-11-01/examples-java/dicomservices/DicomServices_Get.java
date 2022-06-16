import com.azure.core.util.Context;

/** Samples for DicomServices Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/dicomservices/DicomServices_Get.json
     */
    /**
     * Sample code: Get a dicomservice.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void getADicomservice(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.dicomServices().getWithResponse("testRG", "workspace1", "blue", Context.NONE);
    }
}
