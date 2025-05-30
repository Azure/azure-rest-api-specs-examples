
import com.azure.resourcemanager.apimanagement.models.BackendContract;
import com.azure.resourcemanager.apimanagement.models.BackendTlsProperties;

/**
 * Samples for Backend Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementUpdateBackend.json
     */
    /**
     * Sample code: ApiManagementUpdateBackend.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementUpdateBackend(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        BackendContract resource = manager.backends()
            .getWithResponse("rg1", "apimService1", "proxybackend", com.azure.core.util.Context.NONE).getValue();
        resource.update().withDescription("description5308")
            .withTls(new BackendTlsProperties().withValidateCertificateChain(false).withValidateCertificateName(true))
            .withIfMatch("*").apply();
    }
}
