
import com.azure.resourcemanager.apimanagement.models.ExportApi;
import com.azure.resourcemanager.apimanagement.models.ExportFormat;

/**
 * Samples for WorkspaceApiExport Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspaceApiExportInOpenApi2dot0.json
     */
    /**
     * Sample code: ApiManagementGetWorkspaceApiExportInOpenApi2dot0.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetWorkspaceApiExportInOpenApi2dot0(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiExports().getWithResponse("rg1", "apimService1", "wks1", "echo-api",
            ExportFormat.SWAGGER_LINK, ExportApi.TRUE, com.azure.core.util.Context.NONE);
    }
}
