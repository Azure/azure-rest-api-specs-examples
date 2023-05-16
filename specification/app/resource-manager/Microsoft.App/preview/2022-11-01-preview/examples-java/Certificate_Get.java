/** Samples for Certificates Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/Certificate_Get.json
     */
    /**
     * Sample code: Get Certificate.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getCertificate(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .certificates()
            .getWithResponse(
                "examplerg", "testcontainerenv", "certificate-firendly-name", com.azure.core.util.Context.NONE);
    }
}
