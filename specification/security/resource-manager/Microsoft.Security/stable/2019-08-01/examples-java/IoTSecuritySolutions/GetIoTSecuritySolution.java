
/**
 * Samples for IotSecuritySolution GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/
     * GetIoTSecuritySolution.json
     */
    /**
     * Sample code: Get a IoT security solution.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getAIoTSecuritySolution(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.iotSecuritySolutions().getByResourceGroupWithResponse("MyGroup", "default",
            com.azure.core.util.Context.NONE);
    }
}
