/** Samples for IntegrationServiceEnvironments List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_ListBySubscription.json
     */
    /**
     * Sample code: List integration service environments by subscription.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void listIntegrationServiceEnvironmentsBySubscription(
        com.azure.resourcemanager.logic.LogicManager manager) {
        manager.integrationServiceEnvironments().list(null, com.azure.core.util.Context.NONE);
    }
}
