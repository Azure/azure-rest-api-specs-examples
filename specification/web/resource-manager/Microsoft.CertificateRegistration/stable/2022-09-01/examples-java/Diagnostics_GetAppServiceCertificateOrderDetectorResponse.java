/** Samples for CertificateOrdersDiagnostics GetAppServiceCertificateOrderDetectorResponse. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2022-09-01/examples/Diagnostics_GetAppServiceCertificateOrderDetectorResponse.json
     */
    /**
     * Sample code: Get app service certificate order detector response.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAppServiceCertificateOrderDetectorResponse(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getCertificateOrdersDiagnostics()
            .getAppServiceCertificateOrderDetectorResponseWithResponse(
                "Sample-WestUSResourceGroup",
                "SampleCertificateOrderName",
                "AutoRenewStatus",
                null,
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
