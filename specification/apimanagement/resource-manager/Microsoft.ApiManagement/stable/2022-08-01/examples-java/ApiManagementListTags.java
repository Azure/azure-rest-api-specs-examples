/** Samples for Tag ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListTags.json
     */
    /**
     * Sample code: ApiManagementListTags.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListTags(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.tags().listByService("rg1", "apimService1", null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
