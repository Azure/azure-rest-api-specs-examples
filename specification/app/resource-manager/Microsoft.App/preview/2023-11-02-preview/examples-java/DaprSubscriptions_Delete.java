
/**
 * Samples for DaprSubscriptions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/DaprSubscriptions_Delete.
     * json
     */
    /**
     * Sample code: Delete dapr subscription.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deleteDaprSubscription(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.daprSubscriptions().deleteWithResponse("examplerg", "myenvironment", "mysubscription",
            com.azure.core.util.Context.NONE);
    }
}
