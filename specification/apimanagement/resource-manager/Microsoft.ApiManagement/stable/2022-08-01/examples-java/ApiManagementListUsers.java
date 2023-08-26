/** Samples for User ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListUsers.json
     */
    /**
     * Sample code: ApiManagementListUsers.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListUsers(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.users().listByService("rg1", "apimService1", null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
