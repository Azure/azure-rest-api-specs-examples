
/**
 * Samples for AIManagers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/AIManagers_ListBySubscription.json
     */
    /**
     * Sample code: Lists AI Manager resources by subscription.
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void listsAIManagerResourcesBySubscription(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.aIManagers().list(com.azure.core.util.Context.NONE);
    }
}
