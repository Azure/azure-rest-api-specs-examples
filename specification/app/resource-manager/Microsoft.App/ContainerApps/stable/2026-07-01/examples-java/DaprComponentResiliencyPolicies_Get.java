
/**
 * Samples for DaprComponentResiliencyPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/DaprComponentResiliencyPolicies_Get.json
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
