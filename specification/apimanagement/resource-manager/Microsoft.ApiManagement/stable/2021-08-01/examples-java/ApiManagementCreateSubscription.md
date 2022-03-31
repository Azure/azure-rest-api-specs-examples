Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.SubscriptionCreateParameters;

/** Samples for Subscription CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateSubscription.json
     */
    /**
     * Sample code: ApiManagementCreateSubscription.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateSubscription(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .subscriptions()
            .createOrUpdateWithResponse(
                "rg1",
                "apimService1",
                "testsub",
                new SubscriptionCreateParameters()
                    .withOwnerId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/57127d485157a511ace86ae7")
                    .withScope(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/products/5600b59475ff190048060002")
                    .withDisplayName("testsub"),
                null,
                null,
                null,
                Context.NONE);
    }
}
```
