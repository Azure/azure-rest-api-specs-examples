
/**
 * Samples for IotSecuritySolution Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/
     * DeleteIoTSecuritySolution.json
     */
    /**
     * Sample code: Delete an IoT security solution.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteAnIoTSecuritySolution(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.iotSecuritySolutions().deleteByResourceGroupWithResponse("MyGroup", "default",
            com.azure.core.util.Context.NONE);
    }
}
