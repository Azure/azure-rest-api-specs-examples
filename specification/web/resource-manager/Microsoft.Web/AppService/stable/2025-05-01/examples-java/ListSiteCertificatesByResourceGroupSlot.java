
/**
 * Samples for SiteCertificates ListSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListSiteCertificatesByResourceGroupSlot.json
     */
    /**
     * Sample code: List Certificates by resource group for a slot.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        listCertificatesByResourceGroupForASlot(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getSiteCertificates().listSlot("testrg123", "testSiteName", "staging",
            com.azure.core.util.Context.NONE);
    }
}
