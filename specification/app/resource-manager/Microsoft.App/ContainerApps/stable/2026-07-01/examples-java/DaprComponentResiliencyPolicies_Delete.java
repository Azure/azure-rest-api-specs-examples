
/**
 * Samples for DaprComponentResiliencyPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/DaprComponentResiliencyPolicies_Delete.json
     */
    /**
     * Sample code: Delete dapr component resiliency policy.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        deleteDaprComponentResiliencyPolicy(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.daprComponentResiliencyPolicies().deleteWithResponse("examplerg", "myenvironment", "mydaprcomponent",
            "myresiliencypolicy", com.azure.core.util.Context.NONE);
    }
}
