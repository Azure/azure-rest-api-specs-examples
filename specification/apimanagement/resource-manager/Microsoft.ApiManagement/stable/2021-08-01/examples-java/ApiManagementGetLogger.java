import com.azure.core.util.Context;

/** Samples for Logger Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetLogger.json
     */
    /**
     * Sample code: ApiManagementGetLogger.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetLogger(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.loggers().getWithResponse("rg1", "apimService1", "templateLogger", Context.NONE);
    }
}
