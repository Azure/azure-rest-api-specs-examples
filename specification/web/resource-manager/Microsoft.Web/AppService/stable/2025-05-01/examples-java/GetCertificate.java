
/**
 * Samples for Certificates GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/GetCertificate.json
     */
    /**
     * Sample code: Get Certificate.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getCertificate(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getCertificates().getByResourceGroupWithResponse("testrg123", "testc6282",
            com.azure.core.util.Context.NONE);
    }
}
