/** Samples for ContentType Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetContentType.json
     */
    /**
     * Sample code: ApiManagementGetContentType.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetContentType(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.contentTypes().getWithResponse("rg1", "apimService1", "page", com.azure.core.util.Context.NONE);
    }
}
