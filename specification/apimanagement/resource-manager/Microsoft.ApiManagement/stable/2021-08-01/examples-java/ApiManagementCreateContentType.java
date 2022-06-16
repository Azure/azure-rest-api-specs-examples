import com.azure.core.util.Context;

/** Samples for ContentType CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateContentType.json
     */
    /**
     * Sample code: ApiManagementCreateContentType.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateContentType(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.contentTypes().createOrUpdateWithResponse("rg1", "apimService1", "page", null, Context.NONE);
    }
}
