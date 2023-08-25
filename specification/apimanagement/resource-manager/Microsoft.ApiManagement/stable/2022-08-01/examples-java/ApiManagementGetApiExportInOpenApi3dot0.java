import com.azure.resourcemanager.apimanagement.models.ExportApi;
import com.azure.resourcemanager.apimanagement.models.ExportFormat;

/** Samples for ApiExport Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetApiExportInOpenApi3dot0.json
     */
    /**
     * Sample code: ApiManagementGetApiExportInOpenApi3dot0.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetApiExportInOpenApi3dot0(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiExports()
            .getWithResponse(
                "rg1",
                "apimService1",
                "aid9676",
                ExportFormat.OPENAPI_LINK,
                ExportApi.TRUE,
                com.azure.core.util.Context.NONE);
    }
}
