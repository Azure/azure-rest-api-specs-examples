
/**
 * Samples for Certificates ListByIotHub.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/iothub_listcertificates.json
     */
    /**
     * Sample code: Certificates_ListByIotHub.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void certificatesListByIotHub(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.certificates().listByIotHubWithResponse("myResourceGroup", "testhub", com.azure.core.util.Context.NONE);
    }
}
