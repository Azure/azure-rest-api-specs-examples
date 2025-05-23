
import com.azure.resourcemanager.apimanagement.models.LoggerContract;
import com.azure.resourcemanager.apimanagement.models.LoggerType;

/**
 * Samples for Logger Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementUpdateLogger.json
     */
    /**
     * Sample code: ApiManagementUpdateLogger.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateLogger(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        LoggerContract resource = manager.loggers()
            .getWithResponse("rg1", "apimService1", "eh1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withLoggerType(LoggerType.AZURE_EVENT_HUB).withDescription("updating description")
            .withIfMatch("*").apply();
    }
}
