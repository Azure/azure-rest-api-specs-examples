import com.azure.resourcemanager.apimanagement.models.TemplateName;

/** Samples for EmailTemplate GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadEmailTemplate.json
     */
    /**
     * Sample code: ApiManagementHeadEmailTemplate.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadEmailTemplate(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .emailTemplates()
            .getEntityTagWithResponse(
                "rg1", "apimService1", TemplateName.NEW_ISSUE_NOTIFICATION_MESSAGE, com.azure.core.util.Context.NONE);
    }
}
