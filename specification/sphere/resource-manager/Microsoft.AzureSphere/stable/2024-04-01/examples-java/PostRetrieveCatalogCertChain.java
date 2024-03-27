
/**
 * Samples for Certificates RetrieveCertChain.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/
     * PostRetrieveCatalogCertChain.json
     */
    /**
     * Sample code: Certificates_RetrieveCertChain.
     * 
     * @param manager Entry point to AzureSphereManager.
     */
    public static void certificatesRetrieveCertChain(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager.certificates().retrieveCertChainWithResponse("MyResourceGroup1", "MyCatalog1", "active",
            com.azure.core.util.Context.NONE);
    }
}
