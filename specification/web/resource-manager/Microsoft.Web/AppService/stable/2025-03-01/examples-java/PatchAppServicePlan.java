
import com.azure.resourcemanager.appservice.models.AppServicePlanPatchResource;

/**
 * Samples for AppServicePlans Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/PatchAppServicePlan.json
     */
    /**
     * Sample code: Patch Service plan.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void patchServicePlan(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServicePlans().updateWithResponse("testrg123", "testsf6141",
            new AppServicePlanPatchResource().withKind("app"), com.azure.core.util.Context.NONE);
    }
}
