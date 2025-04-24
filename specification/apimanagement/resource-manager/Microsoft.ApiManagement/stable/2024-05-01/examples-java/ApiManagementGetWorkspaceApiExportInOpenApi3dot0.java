
import com.azure.resourcemanager.apimanagement.models.ExportApi;
import com.azure.resourcemanager.apimanagement.models.ExportFormat;

/**
 * Samples for WorkspaceApiExport Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetWorkspaceApiExportInOpenApi3dot0.json
     */
    /**
     * Sample code: ApiManagementGetWorkspaceApiExportInOpenApi3dot0.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetWorkspaceApiExportInOpenApi3dot0(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiExports().getWithResponse("rg1", "apimService1", "wks1", "aid9676",
            ExportFormat.OPENAPI_LINK, ExportApi.TRUE, com.azure.core.util.Context.NONE);
    }
}
