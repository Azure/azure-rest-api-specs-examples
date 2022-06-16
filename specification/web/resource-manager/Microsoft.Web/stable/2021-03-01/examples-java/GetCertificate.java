import com.azure.core.util.Context;

/** Samples for Certificates GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GetCertificate.json
     */
    /**
     * Sample code: Get Certificate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getCertificate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getCertificates()
            .getByResourceGroupWithResponse("testrg123", "testc6282", Context.NONE);
    }
}
