
/**
 * Samples for Certificates Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_certificatesdelete.json
     */
    /**
     * Sample code: Certificates_Delete.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void certificatesDelete(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.certificates().deleteWithResponse("myResourceGroup", "myhub", "cert", "AAAAAAAADGk=",
            com.azure.core.util.Context.NONE);
    }
}
