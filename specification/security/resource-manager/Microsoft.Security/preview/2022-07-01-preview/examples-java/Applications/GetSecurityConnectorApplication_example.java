import com.azure.core.util.Context;

/** Samples for SecurityConnectorApplicationOperation Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-07-01-preview/examples/Applications/GetSecurityConnectorApplication_example.json
     */
    /**
     * Sample code: Get security applications by specific applicationId.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getSecurityApplicationsBySpecificApplicationId(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .securityConnectorApplicationOperations()
            .getWithResponse("gcpResourceGroup", "gcpconnector", "ad9a8e26-29d9-4829-bb30-e597a58cdbb8", Context.NONE);
    }
}
