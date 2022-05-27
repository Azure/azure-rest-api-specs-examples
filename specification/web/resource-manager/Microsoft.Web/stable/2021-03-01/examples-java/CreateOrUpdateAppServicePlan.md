Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appservice.fluent.models.AppServicePlanInner;
import com.azure.resourcemanager.appservice.models.SkuDescription;

/** Samples for AppServicePlans CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/CreateOrUpdateAppServicePlan.json
     */
    /**
     * Sample code: Create Or Update App Service plan.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAppServicePlan(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getAppServicePlans()
            .createOrUpdate(
                "testrg123",
                "testsf6141",
                new AppServicePlanInner()
                    .withLocation("East US")
                    .withSku(
                        new SkuDescription()
                            .withName("P1")
                            .withTier("Premium")
                            .withSize("P1")
                            .withFamily("P")
                            .withCapacity(1))
                    .withKind("app"),
                Context.NONE);
    }
}
```
