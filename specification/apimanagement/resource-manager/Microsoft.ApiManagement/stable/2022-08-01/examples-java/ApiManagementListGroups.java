/** Samples for Group ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListGroups.json
     */
    /**
     * Sample code: ApiManagementListGroups.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListGroups(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.groups().listByService("rg1", "apimService1", null, null, null, com.azure.core.util.Context.NONE);
    }
}
