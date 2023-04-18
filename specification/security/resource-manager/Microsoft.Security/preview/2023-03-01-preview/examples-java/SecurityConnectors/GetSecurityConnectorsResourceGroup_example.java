/** Samples for SecurityConnectors ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2023-03-01-preview/examples/SecurityConnectors/GetSecurityConnectorsResourceGroup_example.json
     */
    /**
     * Sample code: List all security connectors of a specified resource group.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void listAllSecurityConnectorsOfASpecifiedResourceGroup(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.securityConnectors().listByResourceGroup("exampleResourceGroup", com.azure.core.util.Context.NONE);
    }
}
