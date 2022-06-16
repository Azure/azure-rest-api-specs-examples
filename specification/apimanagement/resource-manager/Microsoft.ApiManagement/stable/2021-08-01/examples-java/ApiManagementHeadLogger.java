import com.azure.core.util.Context;

/** Samples for Logger GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadLogger.json
     */
    /**
     * Sample code: ApiManagementHeadLogger.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadLogger(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.loggers().getEntityTagWithResponse("rg1", "apimService1", "templateLogger", Context.NONE);
    }
}
