import com.azure.core.util.Context;

/** Samples for Certificates Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_certificatesdelete.json
     */
    /**
     * Sample code: Certificates_Delete.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void certificatesDelete(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.certificates().deleteWithResponse("myResourceGroup", "myhub", "cert", "AAAAAAAADGk=", Context.NONE);
    }
}
