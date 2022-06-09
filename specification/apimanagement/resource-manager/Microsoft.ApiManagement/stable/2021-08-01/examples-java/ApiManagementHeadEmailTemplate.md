```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.TemplateName;

/** Samples for EmailTemplate GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadEmailTemplate.json
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
            .getEntityTagWithResponse("rg1", "apimService1", TemplateName.NEW_ISSUE_NOTIFICATION_MESSAGE, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
