/** Samples for Region ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListRegions.json
     */
    /**
     * Sample code: ApiManagementListRegions.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListRegions(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.regions().listByService("rg1", "apimService1", com.azure.core.util.Context.NONE);
    }
}
