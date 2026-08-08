
/**
 * Samples for AIManagers ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/AIManagers_ListByResourceGroup.json
     */
    /**
     * Sample code: Lists AI Manager resources by resource group.
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void listsAIManagerResourcesByResourceGroup(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.aIManagers().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
