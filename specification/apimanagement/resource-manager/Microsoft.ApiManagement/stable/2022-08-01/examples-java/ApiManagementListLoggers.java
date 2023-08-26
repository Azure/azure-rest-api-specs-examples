/** Samples for Logger ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListLoggers.json
     */
    /**
     * Sample code: ApiManagementListLoggers.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListLoggers(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.loggers().listByService("rg1", "apimService1", null, null, null, com.azure.core.util.Context.NONE);
    }
}
