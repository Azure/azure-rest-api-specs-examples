import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.TemplateName;

/** Samples for EmailTemplate Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetTemplate.json
     */
    /**
     * Sample code: ApiManagementGetTemplate.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetTemplate(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .emailTemplates()
            .getWithResponse("rg1", "apimService1", TemplateName.NEW_ISSUE_NOTIFICATION_MESSAGE, Context.NONE);
    }
}
