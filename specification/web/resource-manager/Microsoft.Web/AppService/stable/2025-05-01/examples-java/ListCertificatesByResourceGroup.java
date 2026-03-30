
/**
 * Samples for Certificates ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListCertificatesByResourceGroup.json
     */
    /**
     * Sample code: List Certificates by resource group.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listCertificatesByResourceGroup(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getCertificates().listByResourceGroup("testrg123", com.azure.core.util.Context.NONE);
    }
}
