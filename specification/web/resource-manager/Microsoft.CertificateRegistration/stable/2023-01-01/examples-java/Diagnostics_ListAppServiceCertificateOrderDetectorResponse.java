
/**
 * Samples for CertificateOrdersDiagnostics ListAppServiceCertificateOrderDetectorResponse.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2023-01-01/examples/
     * Diagnostics_ListAppServiceCertificateOrderDetectorResponse.json
     */
    /**
     * Sample code: List app service certificate detector response.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAppServiceCertificateDetectorResponse(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getCertificateOrdersDiagnostics()
            .listAppServiceCertificateOrderDetectorResponse("Sample-WestUSResourceGroup", "SampleCertificateOrderName",
                com.azure.core.util.Context.NONE);
    }
}
