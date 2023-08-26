import com.azure.resourcemanager.apimanagement.models.TemplateName;

/** Samples for EmailTemplate Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteTemplate.json
     */
    /**
     * Sample code: ApiManagementDeleteTemplate.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteTemplate(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .emailTemplates()
            .deleteWithResponse(
                "rg1",
                "apimService1",
                TemplateName.NEW_ISSUE_NOTIFICATION_MESSAGE,
                "*",
                com.azure.core.util.Context.NONE);
    }
}
