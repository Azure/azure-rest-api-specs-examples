import com.azure.core.util.Context;

/** Samples for GroupUser List. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListGroupUsers.json
     */
    /**
     * Sample code: ApiManagementListGroupUsers.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListGroupUsers(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.groupUsers().list("rg1", "apimService1", "57d2ef278aa04f0888cba3f3", null, null, null, Context.NONE);
    }
}
