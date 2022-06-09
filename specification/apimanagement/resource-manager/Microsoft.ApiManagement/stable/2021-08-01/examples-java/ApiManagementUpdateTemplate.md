```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.EmailTemplateContract;
import com.azure.resourcemanager.apimanagement.models.TemplateName;

/** Samples for EmailTemplate Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateTemplate.json
     */
    /**
     * Sample code: ApiManagementUpdateTemplate.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateTemplate(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        EmailTemplateContract resource =
            manager
                .emailTemplates()
                .getWithResponse("rg1", "apimService1", TemplateName.NEW_ISSUE_NOTIFICATION_MESSAGE, Context.NONE)
                .getValue();
        resource
            .update()
            .withSubject("Your request $IssueName was received")
            .withBody(
                "<!DOCTYPE html >\r\n"
                    + "<html>\r\n"
                    + "  <head />\r\n"
                    + "  <body>\r\n"
                    + "    <p style=\"font-size:12pt;font-family:'Segoe UI'\">Dear $DevFirstName $DevLastName,</p>\r\n"
                    + "    <p style=\"font-size:12pt;font-family:'Segoe UI'\">\r\n"
                    + "          We are happy to let you know that your request to publish the $AppName application in"
                    + " the gallery has been approved. Your application has been published and can be viewed <a"
                    + " href=\"http://$DevPortalUrl/Applications/Details/$AppId\">here</a>.\r\n"
                    + "        </p>\r\n"
                    + "    <p style=\"font-size:12pt;font-family:'Segoe UI'\">Best,</p>\r\n"
                    + "    <p style=\"font-size:12pt;font-family:'Segoe UI'\">The $OrganizationName API Team</p>\r\n"
                    + "  </body>\r\n"
                    + "</html>")
            .withIfMatch("*")
            .apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
