import com.azure.core.util.Context;

/** Samples for ApiOperation Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteApiOperation.json
     */
    /**
     * Sample code: ApiManagementDeleteApiOperation.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteApiOperation(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiOperations()
            .deleteWithResponse(
                "rg1", "apimService1", "57d2ef278aa04f0888cba3f3", "57d2ef278aa04f0ad01d6cdc", "*", Context.NONE);
    }
}
