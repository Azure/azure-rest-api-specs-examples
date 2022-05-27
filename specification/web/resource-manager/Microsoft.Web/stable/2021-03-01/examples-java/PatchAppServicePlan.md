Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appservice.models.AppServicePlanPatchResource;

/** Samples for AppServicePlans Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/PatchAppServicePlan.json
     */
    /**
     * Sample code: Patch Service plan.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void patchServicePlan(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getAppServicePlans()
            .updateWithResponse(
                "testrg123", "testsf6141", new AppServicePlanPatchResource().withKind("app"), Context.NONE);
    }
}
```
