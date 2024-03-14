
/**
 * Samples for SecurityConnectorApplicationOperation Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2022-07-01-preview/examples/Applications/
     * DeleteSecurityConnectorApplication_example.json
     */
    /**
     * Sample code: Delete security Application.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteSecurityApplication(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.securityConnectorApplicationOperations().deleteWithResponse("gcpResourceGroup", "gcpconnector",
            "ad9a8e26-29d9-4829-bb30-e597a58cdbb8", com.azure.core.util.Context.NONE);
    }
}
