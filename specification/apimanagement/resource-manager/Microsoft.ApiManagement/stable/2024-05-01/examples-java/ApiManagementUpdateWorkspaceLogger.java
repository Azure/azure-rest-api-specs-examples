
import com.azure.resourcemanager.apimanagement.models.LoggerType;
import com.azure.resourcemanager.apimanagement.models.LoggerUpdateContract;

/**
 * Samples for WorkspaceLogger Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementUpdateWorkspaceLogger.json
     */
    /**
     * Sample code: ApiManagementUpdateWorkspaceLogger.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementUpdateWorkspaceLogger(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceLoggers().updateWithResponse(
            "rg1", "apimService1", "wks1", "eh1", "*", new LoggerUpdateContract()
                .withLoggerType(LoggerType.AZURE_EVENT_HUB).withDescription("updating description"),
            com.azure.core.util.Context.NONE);
    }
}
