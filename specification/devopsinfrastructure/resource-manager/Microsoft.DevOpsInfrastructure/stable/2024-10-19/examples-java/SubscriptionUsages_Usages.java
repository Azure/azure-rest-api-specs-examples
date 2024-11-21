
/**
 * Samples for SubscriptionUsages Usages.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-10-19/SubscriptionUsages_Usages.json
     */
    /**
     * Sample code: SubscriptionUsages_Usages.
     * 
     * @param manager Entry point to DevOpsInfrastructureManager.
     */
    public static void
        subscriptionUsagesUsages(com.azure.resourcemanager.devopsinfrastructure.DevOpsInfrastructureManager manager) {
        manager.subscriptionUsages().usages("eastus", com.azure.core.util.Context.NONE);
    }
}
