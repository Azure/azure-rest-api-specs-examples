/** Samples for GroupUser Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateGroupUser.json
     */
    /**
     * Sample code: ApiManagementCreateGroupUser.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateGroupUser(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .groupUsers()
            .createWithResponse(
                "rg1", "apimService1", "tempgroup", "59307d350af58404d8a26300", com.azure.core.util.Context.NONE);
    }
}
