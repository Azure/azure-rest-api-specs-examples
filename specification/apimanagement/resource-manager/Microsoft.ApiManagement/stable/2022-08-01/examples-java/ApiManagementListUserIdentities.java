/** Samples for UserIdentities List. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListUserIdentities.json
     */
    /**
     * Sample code: ApiManagementListUserIdentities.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListUserIdentities(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .userIdentities()
            .list("rg1", "apimService1", "57f2af53bb17172280f44057", com.azure.core.util.Context.NONE);
    }
}
