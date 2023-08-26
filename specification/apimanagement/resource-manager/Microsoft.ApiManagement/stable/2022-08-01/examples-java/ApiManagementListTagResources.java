/** Samples for TagResource ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListTagResources.json
     */
    /**
     * Sample code: ApiManagementListTagResources.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListTagResources(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.tagResources().listByService("rg1", "apimService1", null, null, null, com.azure.core.util.Context.NONE);
    }
}
