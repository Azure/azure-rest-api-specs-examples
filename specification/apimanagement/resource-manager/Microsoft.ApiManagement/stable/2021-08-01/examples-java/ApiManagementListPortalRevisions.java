import com.azure.core.util.Context;

/** Samples for PortalRevision ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListPortalRevisions.json
     */
    /**
     * Sample code: ApiManagementListPortalRevisions.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListPortalRevisions(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.portalRevisions().listByService("rg1", "apimService1", null, null, null, Context.NONE);
    }
}
