import com.azure.core.util.Context;

/** Samples for EmailTemplate ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListTemplates.json
     */
    /**
     * Sample code: ApiManagementListTemplates.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListTemplates(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.emailTemplates().listByService("rg1", "apimService1", null, null, null, Context.NONE);
    }
}
