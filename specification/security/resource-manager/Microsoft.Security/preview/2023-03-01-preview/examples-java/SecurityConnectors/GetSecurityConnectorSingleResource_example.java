/** Samples for SecurityConnectors GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2023-03-01-preview/examples/SecurityConnectors/GetSecurityConnectorSingleResource_example.json
     */
    /**
     * Sample code: Retrieve a security connector.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void retrieveASecurityConnector(com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .securityConnectors()
            .getByResourceGroupWithResponse(
                "exampleResourceGroup", "exampleSecurityConnectorName", com.azure.core.util.Context.NONE);
    }
}
