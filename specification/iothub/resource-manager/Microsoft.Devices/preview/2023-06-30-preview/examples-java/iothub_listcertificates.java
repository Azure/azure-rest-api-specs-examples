/** Samples for Certificates ListByIotHub. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2023-06-30-preview/examples/iothub_listcertificates.json
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
