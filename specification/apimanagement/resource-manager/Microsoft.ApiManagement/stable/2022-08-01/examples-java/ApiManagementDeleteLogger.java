/** Samples for Logger Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteLogger.json
     */
    /**
     * Sample code: ApiManagementDeleteLogger.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteLogger(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.loggers().deleteWithResponse("rg1", "apimService1", "loggerId", "*", com.azure.core.util.Context.NONE);
    }
}
