```java
import com.azure.resourcemanager.apimanagement.models.State;
import java.time.OffsetDateTime;

/** Samples for ApiIssue CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateApiIssue.json
     */
    /**
     * Sample code: ApiManagementCreateApiIssue.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateApiIssue(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiIssues()
            .define("57d2ef278aa04f0ad01d6cdc")
            .withExistingApi("rg1", "apimService1", "57d1f7558aa04f15146d9d8a")
            .withTitle("New API issue")
            .withDescription("New API issue description")
            .withUserId(
                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/1")
            .withCreatedDate(OffsetDateTime.parse("2018-02-01T22:21:20.467Z"))
            .withState(State.OPEN)
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
