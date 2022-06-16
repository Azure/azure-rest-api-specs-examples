import com.azure.core.util.Context;

/** Samples for Certificates ListByIotHub. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_listcertificates.json
     */
    /**
     * Sample code: Certificates_ListByIotHub.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void certificatesListByIotHub(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.certificates().listByIotHubWithResponse("myResourceGroup", "testhub", Context.NONE);
    }
}
