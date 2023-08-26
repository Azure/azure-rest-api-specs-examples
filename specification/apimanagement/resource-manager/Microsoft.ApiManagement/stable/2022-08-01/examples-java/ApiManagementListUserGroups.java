/** Samples for UserGroup List. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListUserGroups.json
     */
    /**
     * Sample code: ApiManagementListUserGroups.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListUserGroups(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .userGroups()
            .list(
                "rg1", "apimService1", "57681833a40f7eb6c49f6acf", null, null, null, com.azure.core.util.Context.NONE);
    }
}
