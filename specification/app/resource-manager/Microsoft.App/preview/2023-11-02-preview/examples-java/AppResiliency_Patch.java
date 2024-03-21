
import com.azure.resourcemanager.appcontainers.models.AppResiliency;
import com.azure.resourcemanager.appcontainers.models.TimeoutPolicy;

/**
 * Samples for AppResiliency Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/AppResiliency_Patch.json
     */
    /**
     * Sample code: Update App Resiliency.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void updateAppResiliency(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        AppResiliency resource = manager.appResiliencies()
            .getWithResponse("rg", "testcontainerApp0", "resiliency-policy-1", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update()
            .withTimeoutPolicy(new TimeoutPolicy().withResponseTimeoutInSeconds(30).withConnectionTimeoutInSeconds(40))
            .apply();
    }
}
