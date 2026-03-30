
/**
 * Samples for Certificates List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListCertificates.json
     */
    /**
     * Sample code: List Certificates for subscription.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listCertificatesForSubscription(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getCertificates().list(null, com.azure.core.util.Context.NONE);
    }
}
