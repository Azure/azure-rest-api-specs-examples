
/** Samples for GlobalAdministrator ElevateAccess. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/authorization/resource-manager/Microsoft.Authorization/stable/2015-07-01/examples/ElevateAccess.
     * json
     */
    /**
     * Sample code: Elevate access global administrator.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void elevateAccessGlobalAdministrator(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.accessManagement().roleAssignments().manager().roleServiceClient().getGlobalAdministrators()
            .elevateAccessWithResponse(com.azure.core.util.Context.NONE);
    }
}
