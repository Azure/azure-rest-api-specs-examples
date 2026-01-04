
import com.azure.resourcemanager.appservice.fluent.models.AppServicePlanInner;
import com.azure.resourcemanager.appservice.models.SkuDescription;

/**
 * Samples for AppServicePlans CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * CreateOrUpdateAppServicePlan.json
     */
    /**
     * Sample code: Create Or Update App Service plan.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAppServicePlan(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServicePlans().createOrUpdate("testrg123", "testsf6141",
            new AppServicePlanInner().withLocation("East US").withSku(
                new SkuDescription().withName("P1").withTier("Premium").withSize("P1").withFamily("P").withCapacity(1))
                .withKind("app"),
            com.azure.core.util.Context.NONE);
    }
}
