/** Samples for ApiSchema ListByApi. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListApiSchemas.json
     */
    /**
     * Sample code: ApiManagementListApiSchemas.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListApiSchemas(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiSchemas()
            .listByApi(
                "rg1", "apimService1", "59d5b28d1f7fab116c282650", null, null, null, com.azure.core.util.Context.NONE);
    }
}
