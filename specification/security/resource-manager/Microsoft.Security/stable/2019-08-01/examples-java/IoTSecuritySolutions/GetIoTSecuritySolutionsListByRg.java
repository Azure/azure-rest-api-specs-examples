import com.azure.core.util.Context;

/** Samples for IotSecuritySolution ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/GetIoTSecuritySolutionsListByRg.json
     */
    /**
     * Sample code: List IoT Security solutions by resource group.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void listIoTSecuritySolutionsByResourceGroup(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.iotSecuritySolutions().listByResourceGroup("MyGroup", null, Context.NONE);
    }
}
