import com.azure.core.util.Context;

/** Samples for Certificates Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/Certificate_Get.json
     */
    /**
     * Sample code: Get Certificate.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getCertificate(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .certificates()
            .getWithResponse("examplerg", "testcontainerenv", "certificate-firendly-name", Context.NONE);
    }
}
