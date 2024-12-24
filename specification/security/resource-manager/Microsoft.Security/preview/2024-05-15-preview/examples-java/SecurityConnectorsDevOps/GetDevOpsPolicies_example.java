
/**
 * Samples for DevOpsPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2024-05-15-preview/examples/
     * SecurityConnectorsDevOps/GetDevOpsPolicies_example.json
     */
    /**
     * Sample code: Get_DevOpsPolicies.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getDevOpsPolicies(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.devOpsPolicies().getWithResponse("myRg", "mySecurityConnectorName", "myDevOpsPolicy",
            com.azure.core.util.Context.NONE);
    }
}
