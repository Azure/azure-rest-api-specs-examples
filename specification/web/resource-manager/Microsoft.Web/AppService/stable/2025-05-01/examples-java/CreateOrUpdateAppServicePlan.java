
import com.azure.resourcemanager.appservice.fluent.models.AppServicePlanInner;
import com.azure.resourcemanager.appservice.models.SkuDescription;

/**
 * Samples for AppServicePlans CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/CreateOrUpdateAppServicePlan.json
     */
    /**
     * Sample code: Create Or Update App Service plan.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void createOrUpdateAppServicePlan(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServicePlans().createOrUpdate("testrg123", "testsf6141",
            new AppServicePlanInner().withLocation("East US").withSku(
                new SkuDescription().withName("P1").withTier("Premium").withSize("P1").withFamily("P").withCapacity(1))
                .withKind("app"),
            com.azure.core.util.Context.NONE);
    }
}
