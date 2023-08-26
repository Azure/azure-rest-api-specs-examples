/** Samples for ContentType Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteContentType.json
     */
    /**
     * Sample code: ApiManagementDeleteContentType.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteContentType(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.contentTypes().deleteWithResponse("rg1", "apimService1", "page", "*", com.azure.core.util.Context.NONE);
    }
}
