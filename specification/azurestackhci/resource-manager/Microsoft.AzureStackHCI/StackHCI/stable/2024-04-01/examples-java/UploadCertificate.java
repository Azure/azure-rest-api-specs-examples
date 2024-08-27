
import com.azure.resourcemanager.azurestackhci.models.RawCertificateData;
import com.azure.resourcemanager.azurestackhci.models.UploadCertificateRequest;
import java.util.Arrays;

/**
 * Samples for Clusters UploadCertificate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * UploadCertificate.json
     */
    /**
     * Sample code: Upload certificate.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void uploadCertificate(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.clusters().uploadCertificate("test-rg", "myCluster",
            new UploadCertificateRequest()
                .withProperties(new RawCertificateData().withCertificates(Arrays.asList("base64cert", "base64cert"))),
            com.azure.core.util.Context.NONE);
    }
}
