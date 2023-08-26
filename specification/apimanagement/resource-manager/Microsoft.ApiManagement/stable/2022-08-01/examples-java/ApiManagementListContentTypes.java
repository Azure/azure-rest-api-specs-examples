/** Samples for ContentType ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListContentTypes.json
     */
    /**
     * Sample code: ApiManagementListContentTypes.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListContentTypes(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.contentTypes().listByService("rg1", "apimService1", com.azure.core.util.Context.NONE);
    }
}
