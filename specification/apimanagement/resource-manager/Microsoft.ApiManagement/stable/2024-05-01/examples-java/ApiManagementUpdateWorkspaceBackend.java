
import com.azure.resourcemanager.apimanagement.models.BackendTlsProperties;
import com.azure.resourcemanager.apimanagement.models.BackendUpdateParameters;

/**
 * Samples for WorkspaceBackend Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementUpdateWorkspaceBackend.json
     */
    /**
     * Sample code: ApiManagementUpdateWorkspaceBackend.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementUpdateWorkspaceBackend(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceBackends().updateWithResponse("rg1", "apimService1", "wks1", "proxybackend", "*",
            new BackendUpdateParameters().withDescription("description5308").withTls(
                new BackendTlsProperties().withValidateCertificateChain(false).withValidateCertificateName(true)),
            com.azure.core.util.Context.NONE);
    }
}
