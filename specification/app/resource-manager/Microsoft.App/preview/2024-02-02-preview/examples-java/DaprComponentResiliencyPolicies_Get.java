
/**
 * Samples for DaprComponentResiliencyPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * DaprComponentResiliencyPolicies_Get.json
     */
    /**
     * Sample code: Get Dapr component resiliency policy.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        getDaprComponentResiliencyPolicy(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.daprComponentResiliencyPolicies().getWithResponse("examplerg", "myenvironment", "mydaprcomponent",
            "myresiliencypolicy", com.azure.core.util.Context.NONE);
    }
}
