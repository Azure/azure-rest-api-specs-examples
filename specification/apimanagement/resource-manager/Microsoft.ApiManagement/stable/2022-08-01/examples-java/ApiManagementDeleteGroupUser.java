/** Samples for GroupUser Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteGroupUser.json
     */
    /**
     * Sample code: ApiManagementDeleteGroupUser.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteGroupUser(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .groupUsers()
            .deleteWithResponse(
                "rg1", "apimService1", "templategroup", "59307d350af58404d8a26300", com.azure.core.util.Context.NONE);
    }
}
