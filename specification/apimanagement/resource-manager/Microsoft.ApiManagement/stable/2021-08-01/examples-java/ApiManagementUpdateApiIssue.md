```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.IssueContract;
import com.azure.resourcemanager.apimanagement.models.State;

/** Samples for ApiIssue Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateApiIssue.json
     */
    /**
     * Sample code: ApiManagementUpdateApiIssue.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateApiIssue(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        IssueContract resource =
            manager
                .apiIssues()
                .getWithResponse(
                    "rg1", "apimService1", "57d1f7558aa04f15146d9d8a", "57d2ef278aa04f0ad01d6cdc", null, Context.NONE)
                .getValue();
        resource.update().withState(State.CLOSED).withIfMatch("*").apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
