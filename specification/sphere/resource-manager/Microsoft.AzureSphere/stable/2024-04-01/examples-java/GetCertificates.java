
/**
 * Samples for Certificates ListByCatalog.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/GetCertificates.json
     */
    /**
     * Sample code: Certificates_ListByCatalog.
     * 
     * @param manager Entry point to AzureSphereManager.
     */
    public static void certificatesListByCatalog(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager.certificates().listByCatalog("MyResourceGroup1", "MyCatalog1", null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
