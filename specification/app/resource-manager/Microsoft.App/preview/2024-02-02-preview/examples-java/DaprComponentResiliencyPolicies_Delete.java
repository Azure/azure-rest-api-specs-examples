
/**
 * Samples for DaprComponentResiliencyPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * DaprComponentResiliencyPolicies_Delete.json
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
