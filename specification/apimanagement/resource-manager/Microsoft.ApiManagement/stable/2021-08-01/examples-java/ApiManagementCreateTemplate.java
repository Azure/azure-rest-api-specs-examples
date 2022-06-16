import com.azure.resourcemanager.apimanagement.models.TemplateName;

/** Samples for EmailTemplate CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateTemplate.json
     */
    /**
     * Sample code: ApiManagementCreateTemplate.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateTemplate(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .emailTemplates()
            .define(TemplateName.NEW_ISSUE_NOTIFICATION_MESSAGE)
            .withExistingService("rg1", "apimService1")
            .withSubject("Your request for $IssueName was successfully received.")
            .create();
    }
}
