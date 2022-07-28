import com.azure.core.util.Context;

/** Samples for Certificates Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/DeleteCertificate.json
     */
    /**
     * Sample code: Delete Certificate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteCertificate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getCertificates()
            .deleteWithResponse("testrg123", "testc6282", Context.NONE);
    }
}
