
import com.azure.resourcemanager.appservice.models.AppServicePlanPatchResource;

/**
 * Samples for AppServicePlans Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/PatchAppServicePlan.json
     */
    /**
     * Sample code: Patch Service plan.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void patchServicePlan(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServicePlans().updateWithResponse("testrg123", "testsf6141",
            new AppServicePlanPatchResource().withKind("app"), com.azure.core.util.Context.NONE);
    }
}
