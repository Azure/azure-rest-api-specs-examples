
/**
 * Samples for Certificates Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/GetCertificate.json
     */
    /**
     * Sample code: Certificates_Get.
     * 
     * @param manager Entry point to AzureSphereManager.
     */
    public static void certificatesGet(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager.certificates().getWithResponse("MyResourceGroup1", "MyCatalog1", "default",
            com.azure.core.util.Context.NONE);
    }
}
