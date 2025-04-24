
import com.azure.resourcemanager.apimanagement.models.CertificateCreateOrUpdateParameters;

/**
 * Samples for WorkspaceCertificate CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateWorkspaceCertificate.json
     */
    /**
     * Sample code: ApiManagementCreateWorkspaceCertificate.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateWorkspaceCertificate(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceCertificates().createOrUpdateWithResponse("rg1", "apimService1", "wks1", "tempcert",
            new CertificateCreateOrUpdateParameters()
                .withData("****************Base 64 Encoded Certificate *******************************")
                .withPassword("fakeTokenPlaceholder"),
            null, com.azure.core.util.Context.NONE);
    }
}
