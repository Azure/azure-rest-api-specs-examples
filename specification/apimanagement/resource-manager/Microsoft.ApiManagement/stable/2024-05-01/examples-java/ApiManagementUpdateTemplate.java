
import com.azure.resourcemanager.apimanagement.models.EmailTemplateContract;
import com.azure.resourcemanager.apimanagement.models.TemplateName;

/**
 * Samples for EmailTemplate Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementUpdateTemplate.json
     */
    /**
     * Sample code: ApiManagementUpdateTemplate.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementUpdateTemplate(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        EmailTemplateContract resource = manager.emailTemplates().getWithResponse("rg1", "apimService1",
            TemplateName.NEW_ISSUE_NOTIFICATION_MESSAGE, com.azure.core.util.Context.NONE).getValue();
        resource.update().withSubject("Your request $IssueName was received").withBody(
            "<!DOCTYPE html >\r\n<html>\r\n  <head />\r\n  <body>\r\n    <p style=\"font-size:12pt;font-family:'Segoe UI'\">Dear $DevFirstName $DevLastName,</p>\r\n    <p style=\"font-size:12pt;font-family:'Segoe UI'\">\r\n          We are happy to let you know that your request to publish the $AppName application in the gallery has been approved. Your application has been published and can be viewed <a href=\"http://$DevPortalUrl/Applications/Details/$AppId\">here</a>.\r\n        </p>\r\n    <p style=\"font-size:12pt;font-family:'Segoe UI'\">Best,</p>\r\n    <p style=\"font-size:12pt;font-family:'Segoe UI'\">The $OrganizationName API Team</p>\r\n  </body>\r\n</html>")
            .withIfMatch("*").apply();
    }
}
